# Poetry 诗人游戏
A Solidity Smart Contract.
小明发币记教程示例合约。

## 结构

- Owned 持有人合约--用于继承
- ERC20Token ERC20TOKEN标准合约--用于继承
- XmbToken 小明币游戏用代币
- Poetry 游戏主要的规则和操作

## 测试
可以使用[Maian](https://github.com/MAIAN-tool/MAIAN)来测试合约的安全漏洞

### MAIAN使用
- 首先需要安装依赖
1. Go Ethereum, check https://ethereum.github.io/go-ethereum/install/
2. Solidity compiler, check http://solidity.readthedocs.io/en/develop/installing-solidity.html
3. Z3 Theorem prover, check https://github.com/Z3Prover/z3 下载源码编译后作为python扩展安装
4. web3, try pip install web3
5. PyQt5 (only for GUI Maian), try sudo apt install python-pyqt5

### 利用Go与合约的交互

- 需要安装
1. Go Ethereum
2. Solidity complier

我们需要使用go-ethereum中的abigen这个工具，安装方式参考https://github.com/ethereum/go-ethereum

（*make all安装完毕后请在build/bin中寻找）