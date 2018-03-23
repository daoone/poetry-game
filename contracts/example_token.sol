pragma solidity ^0.4.18;

contract Token {
    // token 持有者
    address public owner;
    // token 初始化数量
    uint256 public initSupply;
    // 记录账户余额
    mapping (address => uint256) public balances;
    // 一个事件记录
    event Transfer(address indexed To, uint256 Value);
    // 构造函数
    function Token(uint256 _initSupply) public {
        owner = msg.sender;
        initSupply = _initSupply;
        balances[owner] = initSupply;
    }

    function transfer(address _to, uint256 _value) public returns (bool) {
        if (balances[msg.sender] >= _value && balances[_to] + _value >= balances[_to]) {
            balances[msg.sender] -= _value;
            balances[_to] += _value;
            Transfer(msg.sender, _to, _value);
            return true;
        }
        return false;
    }
}

contract MyWallet {
    Token public mt;
    address public owner;
    uint256 public exchangeRate = 100; // 一个eth等值于100个MyToken

    event ByMyToken(address indexed From, uint256 );

    function MyWallet(uint256 _initSupply) public {
        mt = new Token(_initSupply);
    }

    function buyToken() payable public returns (bool) {
        require(msg.value * 100 < mt.balances[address(this)]);
        uint256 distribution;
        distribution = msg.value * 100;
        mt.balances[msg.sender] += distribution;
        mt.balances[address(this)] -= distribution;
    }
    // 提现
    function withdraw(uint256 _value) public returns (bool) {
        require(_value < this.balance && msg.sender == owner);
        owner.transfer(_value);
    }
}