#include "stack.h"
#include "node.h"
#include <iostream>
#include <stdlib.h>
using namespace std;

Stack createStack() {
    NodePtr head = (NodePtr)malloc(sizeof(Node));
    head->next = NULL;
    return head;
}

void push(Stack s,ElementType x) {
    NodePtr temp = (NodePtr)malloc(sizeof(Node));
    temp->element = x;
    temp->next = s->next;
    s->next = temp;
    
}
ElementType pop(Stack s){
    ElementType temp;
    if(s->next) {
        temp = s->next->element;
        s->next =s->next->next;
        free(s->next);
    }
    return temp;
}

ElementType top(Stack s) {
    if(s->next) {
        return s->next->element;
    } else {
        return 0;
    }
}

void printStack(Stack s) {
    NodePtr temp = s->next;
    while(temp) {
        cout << "element:" << temp->element << endl;
        temp = temp->next;
    }
}
