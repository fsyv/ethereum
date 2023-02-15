// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.7.0 <0.9.0;

import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "@openzeppelin/contracts/utils/math/SignedMath.sol";
import "@openzeppelin/contracts/utils/structs/EnumerableMap.sol";

/**
 * @title 信誉管理的智能合约
 */
contract Reputation {
    /**
     * @dev 获取节点信誉值
     * @param _address 节点地址
     * @return 返回查询节点的信誉值
     */
    function getReputation(address _address) public view returns (uint256) {
        if (reputations.contains(_address)) {
            return reputations.get(_address);
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

        for (uint256 i = 0; i < reputations.length(); i++) {
            // 统计信誉值大于0的节点
            (, uint256 reputation) = reputations.at(i);

            if (reputation > 0) {
                rep_valid_sum += reputation;
                signer_valid_count++;
            }
        }

        if (0 == signer_valid_count) return 0;

        return rep_valid_sum / signer_valid_count;
    }

    /**
     * @dev 更新节点信誉值
     * @param _address 节点地址
     * @param _score 得分
     */
    function updateReputation(address _address, int _score) public {
        // 判断系统中是否已经有该节点
        if (reputations.contains(_address)) {
            uint256 rep = reputations.get(_address);

            if (_score < 0) {
                SafeMath.sub(rep, SignedMath.abs(_score));
            } else {
                SafeMath.add(rep, uint(_score));
            }

            reputations.set(_address, rep);
        } else {
            // 如果地址不存在，则插入
            reputations.set(_address, uint(_score));
        }
    }

    // Add the library methods
    using EnumerableMap for EnumerableMap.AddressToUintMap;
    /**
     * @dev 记录节点信誉值
     *      写入: 不存在则插入
     *      读取: 不存在则为0
     */
    EnumerableMap.AddressToUintMap private reputations;
}
