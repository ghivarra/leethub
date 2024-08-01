class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer
     */
    function majorityElement($nums) {
        $res  = [];

        foreach ($nums as $key => $num)
        {
            if (!isset($res[$num]))
            {
                $res[$num] = 1;

            } else {

                $res[$num]++;
            }
        }

        $key = array_keys($res, max($res));

        return $key[0];
    }
}