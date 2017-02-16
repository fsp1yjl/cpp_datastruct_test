#include <iostream>
#include <cmath>
using namespace std;
int main() {
    int a,b,c,d,e;
    int num;
    cout << "please input a num with length of 5:" <<endl;
    cin >> num;
    a = num%10;
    num /= 10;
    b = num%10;
    num /= 10;
    c = num%10;
    num /= 10;
    d = num%10;
    num /= 10;
    e = num%10;
  
    cout << "the num  is :" << ((((a*10+b)*10 + c)*10 +d)*10 + e) <<endl;
    return 0;
}