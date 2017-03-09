
// make command: g++ list_test.cpp list.cpp -o list

#include "list.h"
#include <iostream>
using namespace std;
int main() {
    List head = createList();
    cout << "head:" << head << endl;
    PtrToNode last = findLast(head);
    cout << "last:" << last << endl;
    int element = 5;
    insertHead(head,element);
    insertHead(head,element+1);
    insertTail(head,3);
    insertHead(head,8);
    int empty_flag = IsEmpty(head);
    cout << "empty_flag:" << empty_flag <<endl;
    // showList(head);
    // deleteElement(head,8);
    // showList(head);
    insertTail(head,8);
    showList(head);
    deleteElement(head,8);
    showList(head);
    deleteElement(head,8);
    showList(head);






    // List list = (List)malloc(sizeof(Node));
    // int a = 21;
    // list->element = a;
    // list->next = NULL;
    // cout << list->element << endl;
    // //list->element = 2;
    // // int is_empty = 0;
    // int is_empty;
    // is_empty = IsEmpty(list);
    // // //cout << "list:" << list->element << endl;
    // cout << "is_empty:" << is_empty << endl;
    return 0;
}

