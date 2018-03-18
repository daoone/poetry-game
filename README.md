# Poetry 诗人游戏 :dancer:
> A Solidity Smart Contract 小明发币记教程示例合约。

目的在于通过一个简单的游戏示例，帮助大家入门以太坊智能合约的开发和部署。

## 结构

- Owned 持有人合约--用于继承
- ERC20Token ERC20TOKEN标准合约--用于继承
- XmbToken 小明币游戏用代币
- Poetry 游戏主要的规则和操作

## 游戏规则

小明拿出了毕生的家当，为了奖励那些对他甜言蜜语的人，以此来填补自己内心的空虚。
因此他想出了这样一个游戏：
- 每个诗人可以写一句诗放到公链上。
- 所有投票人可以对链上的诗进行投票，投票会消耗XMB，消耗1xm XMB代表增加1权重。权重是最后用来判断游戏结果的标准。每个投票人对一首诗只能投票一次。
- 这个游戏为期一个月，一个月后，小明会拿出15eth奖励得分最高的诗的诗人，85eth奖励诗对应的投票人。

## 测试部署
可以使用[Maian](https://github.com/MAIAN-tool/MAIAN)来测试合约的安全漏洞

### MAIAN使用

- 首先需要安装依赖
1. Go Ethereum, check https://ethereum.github.io/go-ethereum/install/
2. Solidity compiler, check http://solidity.readthedocs.io/en/develop/installing-solidity.html
3. Z3 Theorem prover, check https://github.com/Z3Prover/z3 下载源码编译后作为python扩展安装
4. web3, try pip install web3
5. PyQt5 (only for GUI Maian), try sudo apt install python-pyqt5

在tools目录下直接运行
> python maian.py -s example_contracts/Poetry.sol Poetry -c 0

### 利用Go与合约的交互

- 需要安装
1. Golang 1.8+  https://golang.org/
2. Go Ethereum
3. Solidity complier

我们需要使用go-ethereum中的`abigen`这个工具，安装方式参考https://github.com/ethereum/go-ethereum

（*`make all`安装完毕后请在`build/bin`中寻找）

我们使用go-ethereum自带的工具`abigen`生成合约对象脚本
> abigen --sol=contracts/poetry.sol --pkg=main --out=scripts/poetry.go

生成后的go文件包含了我们合约中的所有对外方法。