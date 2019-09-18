package main;

public class Test {

    public  static void main(String[] args){
        String str= "t1y1u1";
        String[] staArr=str.split("1");
        for (String s:staArr){
            System.err.println(s);
        }
    }
}
