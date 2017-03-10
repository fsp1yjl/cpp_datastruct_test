
#include "stack.h"
#include <iostream>
using namespace std;

int main() {
    Stack s = createStack();
    push(s,10);
    push(s,11);
    push(s,12);
    ElementType te = pop(s);
    printStack(s);
    cout << "te::" << te << endl;

    return 0;
}
