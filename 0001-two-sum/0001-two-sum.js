/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function(nums, target) {
    let twoKeys = []

    for (i = 0; i < nums.length; i++) {
        var srcNum = target - nums[i]
        var index = nums.indexOf(srcNum)

        if (index > -1 && index !== i) {
            twoKeys.push(i, index)
            break
        }
    }

    return twoKeys
};