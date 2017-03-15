

typedef char  ElementType;
#include <iostream>
#include <stack>
#include <vector>
#include <string>
#include "binary_tree.h"



using namespace std;
void print_stack(binary_tree_node node);  //中序遍历表达式树
int main() {
    cout << "start test:" << endl;
    char s1[] = {'a','b','+','c','d','e','+','*','*'};
    stack<Binary_tree_node> tree_stack;
    
    vector<char> expr_str(&s1[0],&s1[0]+9);
    cout << "tree size:" << tree_stack.size() << endl;
    cout << "vector size:" << expr_str.size() << endl;
    int i;
    int index;
    
    for( i = 0; i < expr_str.size(); i++) {
        bool is_expr = false;
        cout << "expr_str[" << i << "] =  " << expr_str[i] << "   ********" ;
        binary_tree_node temp_node;
        cout << "ptr:" << &temp_node << endl;
        temp_node.element = expr_str[i];
        temp_node.left = NULL;
        temp_node.right = NULL;
        if( expr_str[i] == '+'|| expr_str[i] == '-' ||expr_str[i] == '*'){ //目前无法处理除法
            is_expr = true;
        }
        if(!is_expr) {
            tree_stack.push(temp_node);
        } else {
             binary_tree_node temp1,temp2;
             PtrToNode left,right;
             if(!tree_stack.empty()) {
                 temp1 = tree_stack.top();
                 //每次出栈的时候需要讲出栈的结构体放入堆中，这样才能保证后续的指针不失效
                 right = new Binary_tree_node(temp1.element,temp1.left,temp1.right);
                //  cout << "temp1->ele:" << temp1.element << endl;
                 tree_stack.pop();
             }
             if(!tree_stack.empty()) {
                 temp2 = tree_stack.top();
                 left = new Binary_tree_node(temp2.element,temp2.left,temp2.right);
                //   cout << "temp2->ele:" << temp2.element << endl;
                 tree_stack.pop();
             }
           
           
            temp_node.right = right;
            temp_node.left = left;
            tree_stack.push(temp_node);
        }
    }
     binary_tree_node node = tree_stack.top();
     print_stack(tree_stack.top());
     cout << endl;
    return 0;
}

void print_stack(binary_tree_node node) {
    //cout << "element--------->" << node.element << endl;
    if(node.left != NULL) {
        print_stack(*(node.left));
    }
    cout << node.element;
    if(node.right != NULL) {
        print_stack(*(node.right));
    }
}
