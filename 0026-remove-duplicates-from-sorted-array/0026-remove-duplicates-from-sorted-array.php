class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer
     */
    function removeDuplicates(&$nums) {
        $nums = array_unique($nums);
        sort($nums);

        return count($nums);
    }
}