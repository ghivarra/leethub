class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer
     */
    function removeDuplicates(&$nums) {
        $max  = 2;
        $inst = 0;
        $last = '';

        foreach ($nums as $key => $item)
        {
            if ($last === $item)
            {
                $inst++;

            } else {

                $inst = 1;
            }

            if ($inst > $max)
            {
                unset($nums[$key]);
            }

            $last = $item;
        }

        return count($nums);
    }
}