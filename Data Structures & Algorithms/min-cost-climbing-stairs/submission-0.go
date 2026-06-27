func minCostClimbingStairs(cost []int) int {
    n := len(cost)
    cache := make([]int, n+1)
    cache[n] = 0;
    cache[n-1] = cost[n-1]
    for i:=n-2; i>=0; i-- {
        cache[i] = cost[i] + min(cache[i+1], cache[i+2])
    }
    return min(cache[0], cache[1])
}
