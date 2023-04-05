// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.7.0 <0.9.0;

import "@openzeppelin/contracts/utils/math/SignedMath.sol";
import "@openzeppelin/contracts/utils/structs/EnumerableMap.sol";

/**
 * @title 信誉管理(Reputation Value)的智能合约
 */
contract RepValue {
    /**
     * @dev 获取账户信誉值
     * @param account 节点地址
     * @return 返回查询节点的信誉值
     */
    function getReputation(address account) public view returns (uint256) {
        if (_reputation.contains(account)) {
            return _reputation.get(account);
        } else {
            return 0;
        }
    }

    /**
     * @dev 获取信誉阈值
     *      这里使用求有效节点(信誉值大于0)平均值的方式
     * @return 信誉阈值
     */
    function getRepThreshold() public view returns (uint256) {
        // 有效节点总和
        uint256 rep_valid_sum = 0;
        // 统计有效节点数量
        uint256 signer_valid_count = 0;

        for (uint256 i = 0; i < _reputation.length(); i++) {
            // 活动信誉值
            (, uint256 reputation) = _reputation.at(i);
            // 统计信誉值大于0的节点
            if (reputation > 0) {
                rep_valid_sum += reputation;
                signer_valid_count++;
            }
        }

        // 防止除0
        if (0 == signer_valid_count) return 0;

        return rep_valid_sum / signer_valid_count;
    }

    /**
     * @dev 更新节点信誉值
     * @param account 节点地址
     * @param reward 奖励
     */
    function updateReputation(address account, int reward) public {
        // 查询系统中是否存在改账户
        if (_reputation.contains(account)) {
            uint256 reputation = _reputation.get(account);

            if (reward < 0) {
                reputation -= SignedMath.abs(reward);
            } else {
                reputation += uint(reward);
            }

            _reputation.set(account, reputation);
        } else {
            // 账户不存在，插入数据

            uint256 reputation = 0;
            if (reward > 0) {
                reputation += uint(reward);
            }

            _reputation.set(account, reputation);
        }
    }

    // Add the library methods
    using EnumerableMap for EnumerableMap.AddressToUintMap;
    /**
     * @dev 记录节点信誉值
     *      写入: 不存在则插入
     *      读取: 不存在则为0
     */
    EnumerableMap.AddressToUintMap private _reputation;
}
