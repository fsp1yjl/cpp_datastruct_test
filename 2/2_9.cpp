#include <iostream>
#include <cmath>
using namespace std;
int main() {
    double a,b,c;
    double s,area;
    cout << "please input a,b,c:" <<endl;
    cin >> a >> b >> c;
    s = (a+b+c)/2.0;
    area = sqrt(s*(s-a)*(s-b)*(s-c));
    cout << "the area is :" << area <<endl;
    return 0;
}