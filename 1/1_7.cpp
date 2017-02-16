#include <iostream>
#include <cmath>
using namespace std;
int main() {
    double money,years,rate,sum;
    cout << "please input money years and rate:" <<endl;
    cin >> money >> years >> rate;
    sum = money * pow((1+rate),years);
    cout << "the sum is :" << sum <<endl;
    return 0;
}