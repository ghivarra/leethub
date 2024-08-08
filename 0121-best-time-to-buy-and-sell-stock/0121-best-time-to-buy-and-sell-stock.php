class Solution {

    /**
     * @param Integer[] $prices
     * @return Integer
     */
    function maxProfit($prices) {
        $tmpPrices = $prices;
        rsort($tmpPrices);

        if (empty($prices) OR ($prices === $tmpPrices))
        {
            return 0;
        }

        // search the highest margins
        $margin = 0;
        $last   = null;

        foreach ($prices as $key => $price)
        {
            if (!is_null($last) && $price >= $last)
            {
                continue;
            }

            $slice = array_slice($prices, ($key + 1));

            if (empty($slice))
            {
                break;
            }

            $max  = max($slice);
            $diff = $max - $price;

            if ($diff > $margin)
            {
                $margin = $diff;
            }

            // input to last
            $last = $price;
        }

        // count the diffs
        return $margin;
    }
}