class Solution {

    /**
     * @param Integer[] $prices
     * @return Integer
     */
    function maxProfit($prices) {
        if (empty($prices))
        {
            return 0;
        }

        // search the highest margins
        $margin = 0;
        $last   = null;

        // search for the minimum
        $min    = min($prices);
        $minKey = array_keys($prices, $min);
        $minKey = $minKey[0];

        // search for the maximum after minimum
        $slice = array_slice($prices, ($minKey + 1));
        
        // count margin
        $margin = empty($slice) ? 0 : max($slice) - $min;

        // count the diffs
        return $margin;
    }
}