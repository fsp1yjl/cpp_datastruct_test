#ifndef _Binary_tree_H
#define _Binary_tree_H

typedef struct Binary_tree_node * PtrToNode;


typedef struct Binary_tree_node {
    ElementType element;
    PtrToNode left;
    PtrToNode right;
    
    Binary_tree_node(){
        element = '0';
        left = 0;
        right = 0;
    }
    Binary_tree_node(ElementType x,PtrToNode l,PtrToNode r) {
        element = x;
        left = l;
        right =r;
    }
} binary_tree_node;
#endif