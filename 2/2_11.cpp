#include <iostream>
#include <cmath>
using namespace std;
int main() {
    double a,b;
    char op;
    cout << "please input a,op and b:"<<endl;
    cin >> a >> op >>b;
    switch(op) {
        case '+':
            cout << "result is :" << a + b << endl;
            break;
        case '-':
            cout << "result is :" << a -b <<endl;
            break;
        default :
            cout << "unknown operater:"<<endl;

    }
    return 0;
}