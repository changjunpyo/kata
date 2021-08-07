class Solution {
    public long solution(int price, int money, int count) {
        long sn = (long)count*(long)(count+1)/2*price;
        
        if (sn - money > 0){
            return sn - (long) money;
        }
        return 0;
    }
}
