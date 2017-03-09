
#include "list.h"
#include <stdlib.h>
#include <iostream>
using namespace std;



int IsEmpty(List L) {
    if(L->next) {
        return 0;
    } else {
        return 1;
    }
}

List createList() {
    List head = (List)malloc(sizeof(Node));
    head->next = NULL;
    return head;
}

void deleteElement(List head, ElementType x) {
    PtrToNode ptr = findPrevious(head,x);
    cout << "previous ptr:" << ptr << "next:" << ptr->next<<  endl;
    if(ptr) {
        PtrToNode temp = ptr->next;
        ptr->next = ptr->next->next;
        free(temp);
    } else {
        cout << "the need delete element" << x << "not found!!!" << endl;
    }

}

PtrToNode find(List head, ElementType x) {
    PtrToNode ptr = head;
    while(ptr) {
        if(ptr->element == x) {
            break;
        }
        ptr = ptr->next;
    }
    return ptr;
}

PtrToNode findPrevious(List head,ElementType x) {
    PtrToNode previousPtr = head;
    while(previousPtr) {
        if(previousPtr->next && previousPtr->next->element == x) {
            break;
        }
        previousPtr = previousPtr->next; 
    }
    return previousPtr;
}

//get the pointer of the last linklist node
PtrToNode findLast(List head) {
    PtrToNode ptr = head;
    while(ptr->next) {
        ptr = ptr->next;
    }
    return ptr;
}

//insert a new node at the behind link list head
void insertHead(List head,ElementType x) {
    PtrToNode nodePtr = (PtrToNode)malloc(sizeof(Node));
    nodePtr->element = x;
    nodePtr->next = head->next;
    head->next = nodePtr;
}

// insert a new node at the tail of the linklist
void insertTail(List head,ElementType x) {
    PtrToNode nodePtr = (PtrToNode)malloc(sizeof(Node));
    nodePtr->element = x;
    nodePtr->next = NULL;
    PtrToNode last = findLast(head);
    last->next = nodePtr;
}

void showList(List head) {
    PtrToNode nodePtr;
    nodePtr = head->next;
    cout << "<link list start---------------------->" << endl;
    cout << "head:::" << head << "next:" << head->next << endl;
    while(nodePtr) {
        cout << "element:" << nodePtr->element  <<"ptr:" << nodePtr << "next:" <<nodePtr->next <<endl;
        nodePtr = nodePtr->next;
    }
    cout << "<link list end---------------------->" << endl;
}

