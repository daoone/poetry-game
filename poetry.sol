pragma solidity ^0.4.18;

// 合约归属
contract Owned {
    address public owner;

    function Owned() public {
        owner = msg.sender;
    }
    // 一个函数执行的前置条件
    modifier onlyOwner() {
        require(msg.sender == owner);
        _;
    }
    // 更改合约所有人
    function changeOwner(address newOwner) onlyOwner public {
        owner = newOwner;
    }
}

// 标准ERC20Token 类似于一个抽象类
contract ERC20Token {
    uint256 public totalSupply;
    function balanceOf(address _owner) constant public returns (uint256 balance);
    function transfer(address _to, uint256 _value) public returns (bool success);
    function transferFrom(address _from, address _to, uint256 _value) public returns (bool success);
    function approve(address _spender, uint256 _value) public returns (bool success);
    function allowance(address _owner, address _spender) constant public returns (uint256 remaining);
    event Transfer(address indexed _from, address indexed _to, uint256 _value);
    event Approval(address indexed _owner, address indexed _spender, uint256 _value);
}

// 自己的token 继承ERC20Token, Owned
// 该合约只由Poetry初始化，所以owner为Poetry
contract XmbToken is ERC20Token, Owned {
    using SafeMath for uint256; // 使用我们自己的计算库扩展
    string public name;
    string public symbol;
    uint8 public decimals;

    mapping (address => uint256) public balances;
    mapping (address => mapping (address => uint256)) allowed;

    event Transfer(address indexed from, address indexed to, uint256 value);
    event Approval(address indexed _owner, address indexed _spender, uint256 _value);

    function XmbToken(uint256 initialSupply, uint8 decimalUnits) public {
        name = "XmbToken";
        symbol = "XMB";
        balances[msg.sender] = initialSupply;
        decimals = decimalUnits;
    }

    function transfer(address _to, uint256 _value) public returns (bool success) {
	    if (balances[msg.sender] >= _value && balances[_to] + _value >= balances[_to]) {
            balances[msg.sender] -= _value;
            balances[_to] += _value;
            Transfer(msg.sender, _to, _value);
            return true;
        }
	    return false;
	}

    function transferFrom(address _from, address _to, uint256 _value) public returns (bool success) {
        require(msg.sender == owner);
        if (balances[_from] >= _value && allowed[_from][msg.sender] >= _value && _value > 0) {
            balances[_to] += _value;
            balances[_from] -= _value;
            allowed[_from][msg.sender] -= _value;
            Transfer(_from, _to, _value);
            return true;
        } 
        return false;
    } 

    function balanceOf(address _owner) constant public returns (uint256 balance) {
        return balances[_owner];
    }

    function approve(address _spender, uint256 _value) public returns (bool success) {
        allowed[msg.sender][_spender] = _value;
        Approval(msg.sender, _spender, _value);
        return true;
    }

    function allowance(address _owner, address _spender) constant public returns (uint256 remaining) {
      return allowed[_owner][_spender];
    }

    // 增发XMB的方法
    function additional(uint256 _value) onlyOwner public returns (bool success) {
        if (balances[owner] + _value >= balances[owner]) {
            balances[owner] += _value;
            return true;
        }
        return false;
    }

}

// 诗歌游戏合约逻辑
contract Poetry is Owned {
    using SafeMath for uint256;

    // 避免结构体的嵌套，我使用过嵌套会出现 UnimplementedFeatureError
    struct Poem {
        string content; // 诗歌内容
        uint256 votes; // 获得的token投票
        uint voteCounts; // 获得的投票人次
        mapping (address => bool) voted;
        address poetAddr; 
    }

    XmbToken public xmb;
    bool public gameover = false; // 结束游戏参与期
    uint256 public poemReward = 15 ether; // 诗人最终奖励
    uint256 public voteReward = 85 ether; // 投票人最终奖励
    uint256 public eachVoterReward; // 评价投票人奖励，减少之后的计算
    Poem[] public poems; // 所有诗歌
    uint256 public maxVotes; // 现有最高投票数
    uint[] public winners; // 现有最高票数的诗歌id 
    uint256 public rechargeLimit = 30 ether; // 最高可一次购买 30 ether 的token 增加刷票成本、也可以用approve控制
    uint256 public rechargeRate = 1000; // XMB兑换以太坊 1:1000


    function Poetry(uint256 initialSupply, uint8 decimalUnits) public {
        xmb = new XmbToken(initialSupply, decimalUnits);
    }

    event PoemAdded(address from, uint poemId);
    event PoemVoted(address from, address to, uint poemId, uint256 value);
    event RewardPushlished(address from, address to, uint256 value);
    event RechargeFaith(address indexed to, uint256 ethValue, uint256 distrbution, uint256 refund);

    // 写诗
    function addPoem(string poemContent) public returns (uint poemId) {
        require(gameover == false);
        poemId = poems.length++;

        Poem storage pm = poems[poemId];
        pm.content = poemContent;
        pm.votes = 0;
        pm.voteCounts = 0;
        pm.poetAddr = msg.sender;
        pm.voted[msg.sender] = true;

        PoemAdded(msg.sender, poemId);
    }

    // 投票
    function votePoem(uint poemId, uint256 _value) public {
        require(!gameover);
        require(xmb.balances(msg.sender) > 0 && _value <= xmb.balances(msg.sender));
        Poem storage pm = poems[poemId];
        require(!pm.voted[msg.sender]);
        pm.voted[msg.sender] = true;

        xmb.transferFrom(msg.sender, pm.poetAddr, _value);
        pm.votes += _value;
        pm.voteCounts ++;

        if (pm.votes > maxVotes) {
            maxVotes = pm.votes;
            resetWinner(poemId);
        } else if (pm.votes == maxVotes) {
            winners.push(poemId);
        }
        
        PoemVoted(msg.sender, pm.poetAddr, poemId, _value);
    }

    // 最终奖励赢家
    function reward() public returns (bool) {
        require((msg.sender == owner) && (this.balance > (poemReward + voteReward)));
        uint256 eachPoetReward = poemReward.div(winners.length);
        uint tmpVoteCounter = 0;
        for (uint i = 0; i <= winners.length-1; i++) {
            poems[winners[i]].poetAddr.transfer(eachPoetReward);
            RewardPushlished(this, poems[winners[i]].poetAddr, eachPoetReward);
            tmpVoteCounter += poems[winners[i]].voteCounts;
        }
        eachVoterReward = voteReward.div(tmpVoteCounter);
        gameover = true;
    }

    // 投票人主动来领奖
    function getVoterReward() public {
        require(gameover);
        for (uint i = 0; i <= winners.length-1; i++) {
            if (poems[winners[i]].voted[msg.sender]) {
                msg.sender.transfer(eachVoterReward);
                poems[winners[i]].voted[msg.sender] = false;
                RewardPushlished(this, msg.sender, eachVoterReward);
            }
        }
    }

    // 刷下现有最高分
    function resetWinner(uint poemId) internal {
        if (winners.length > 0) {
            for (uint i = 0; i < winners.length-1; i++) {
                delete winners[i];
            }
        }
        winners.push(poemId);
    }

    // 查询个人XMB余额
    function getBalance(address owner) view public returns (uint256) {
        return xmb.balanceOf(owner);
    }

    // 兑换XMB
    function buyXmb() payable public {
        uint256 distrbution;
        uint256 _refund = 0;
        // 转入金额大于0 并且转换后小于可分发token的数量
        if ((msg.value > 0) && (tokenExchange(msg.value) < xmb.balances(this))) {
            if (msg.value > rechargeLimit) {
                // 超过兑换限制会把多余的退回去
                _refund = msg.value.sub(rechargeLimit);
                msg.sender.transfer(_refund);
                distrbution = tokenExchange(rechargeLimit);
            } else {
                distrbution = tokenExchange(msg.value);
            }
            xmb.transfer(msg.sender, distrbution);
        }
        RechargeFaith(msg.sender, msg.value, distrbution, _refund);
    }

    function tokenExchange(uint256 inputAmount) internal view returns (uint256) {
        return inputAmount.mul(rechargeRate);
    }

    // function sell(uint sellAmount) public {
    //     // 卖出token 暂不支持
    // }
}

// 安全计算方法库
library SafeMath {
    function mul(uint a, uint b) internal pure returns (uint) {
        uint c = a * b;
        assert(a == 0 || c / a == b);
        return c;
    }

    function div(uint a, uint b) internal pure returns (uint) {
        uint c = a / b;
        return c;
    }

    function sub(uint a, uint b) internal pure returns (uint) {
        assert(b <= a);
        return a - b;
    }

    function add(uint a, uint b) internal pure returns (uint) {
        uint c = a + b;
        assert(c >= a);
        return c;
    }
}