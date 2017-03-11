

#ifndef _Stack_h
#define _Stack_h
typedef int ElementType;
typedef struct Node *NodePtr;
typedef NodePtr Stack;

Stack createStack();

void push(Stack s,ElementType x);
ElementType pop(Stack s);

void printStack(Stack s);

ElementType top(Stack s);

#endif