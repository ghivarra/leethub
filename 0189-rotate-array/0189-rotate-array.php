class Solution {

    /**
     * @param Integer[] $nums
     * @param Integer $k
     * @return NULL
     */
    function rotate(&$nums, $k) {
        $limit = count($nums);

        if ($k > $limit)
        {
            $calc = $k / $limit;

            if (!is_float($calc))
            {
                $k = 0;

            } else {

                $div = intval(floor($calc));
                $k   = $k - ($limit * $div);
            }
        }

        foreach ($nums as $key => $item)
        {
            $finalKey = $key + 1 + $k;

            if ($finalKey > $limit)
            {
                $finalKey = $finalKey - $limit;
            }

            $array[$finalKey] = $item;
        }

        $nums = $array;
        ksort($nums);
    }
}