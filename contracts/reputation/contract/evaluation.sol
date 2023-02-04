// SPDX-License-Identifier: UNLICENSED
pragma solidity >=0.7.0 <0.9.0;

// TODO 这个智能合约应该是每个出块都会执行
// 1. 合约上链
// 2. 合约执行
// 3. 信誉值修改(修改reputation.sol中的值)
contract Evaluation {
    constructor(){

    }

    // 评价目标账户
    address account;

    // 得分
    int score;

    // 指向另外一个合约，评价的内容
    address evaluate_event;

    // 时间戳
    uint timestamp;

    // 签名
    address signer;

    // TODO: 可以添加校验着签名
}
