#ifndef _LIST_H
#define _LIST_H
// struct Node;
typedef struct Node *PtrToNode;
typedef PtrToNode List;
typedef PtrToNode Position;

typedef int ElementType;
//此处不能将Node结构体的定义放在list.cpp中，否则会包找不到完整定义
struct Node {
    ElementType element;
    Position next;
};

int IsEmpty(List L);

List createList();

PtrToNode findLast(List head);

void insertHead(List head,ElementType x);

void insertTail(List head,ElementType x);

void showList(List head);

PtrToNode find(List head, ElementType x);
PtrToNode findPrevious(List head,ElementType x);

void deleteElement(List head, ElementType x);



#endif