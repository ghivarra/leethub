class Solution {

    /**
     * @param Integer[] $nums
     * @param Integer $val
     * @return Integer
     */
    function removeElement(&$nums, $val) {
        foreach ($nums as $key => $item)
        {
            if ($item === $val)
            {
                unset($nums[$key]);
            }
        }

        return count($nums);
    }
}