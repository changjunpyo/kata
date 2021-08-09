class Solution {
    public String solution(int[][] scores) {
        String answer = "";
        int n = scores.length;
        for (int j=0; j<n; j++){
            int count = 0;
            int sum = 0;
            int total = n;
            for (int i=0; i<n; i++){
                sum += scores[i][j];
                if (scores[i][j] > scores[j][j]) count++;
                if (scores[i][j] < scores[j][j]) count--;
            }
            if (count == -(n-1) || count == n-1){
                sum -= scores[j][j];   
                total--;
            }
            answer += calculateGrade(sum/total);
        }
        
        return answer;
    }
    public String calculateGrade(int score){
        if (score >= 90) return "A";
        if (score >= 80) return "B";
        if (score >= 70) return "C";
        if (score >= 50) return "D";
        return "F";
    }
}
