pragma solidity >=0.7.0 <0.9.0;

contract Reputation {
    constructor() public {

    }

    // 记录账户信誉值
    // 写入: 不存在则插入
    // 读取: 不存在则为0
    mapping(address => uint) reputations;
    // TODO 是否标记活跃节点？
}