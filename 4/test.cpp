

#include "stdio.h"
#include "stdlib.h"
#include "time.h"
 
/* 定义 */
#define OK 1
#define ERROR 0
#define TRUE 1
#define FALSE 0
 
/* 别名定义  */
typedef int Status;
typedef int ElemType;
 
/* 单链表存储结构 */
typedef struct Node
{
    ElemType data;          /* 数据域 */
    struct Node *next;      /* 指针域 */
}Node;
 
typedef struct Node *LinkList;      /* 定义LinkList */
 
/* 单链表初始化 */
Status InitList(LinkList *L)
{
    *L = (LinkList)malloc(sizeof(Node));        /* 产生头结点，并使L指向次头结点 */
    if (!(*L))      /* 存储分配失败 */
    {
        return ERROR;
    }
    (*L)->next=NULL;     /* 指针域为空 */
 
    return OK;
}
 
 
/* 初始条件：单链表已存在。操作结果：返回L中数据元素个数 */
int ListLength(LinkList L)
{
    int i = 0;
    LinkList p = L->next;        /* p指向第一个结点 */
    while (p)
    {
        i++;
        p=p->next;
    }
 
    return i;
}
 
 
/* 初始条件：单链表L已存在， 1<= i <= ListLength(L) */
/* 操作结果：在L中第i个位置之前插入新的数据元素e，L的长度加1 */
Status ListInsert(LinkList *L, int i, ElemType e)
{
    int j;
    LinkList p, s;
    p = *L;
    j = 1;
    while (p && j < i)       /* 寻找第i个结点 */
    {
        p = p->next;
        ++j;
    }
    if (!p || j > i)     /* 第i个元素不存在 */
    {
        return ERROR;
    }
 
    s = (LinkList)malloc(sizeof(Node));     /* 生成新结点(C语言标准函数) */
    s->data=e;
    s->next=p->next;      /* 将p的后继结点赋值给s的后继结点 */
    p->next=s;               /* 将s赋值给p的后继结点 */
 
    return OK;
}
 
/* 遍历链表时显示函数 */
Status visit(ElemType c)
{
    printf("%d",c);
    return OK;
}
 
 
/* 初始条件：单链表L已存在 */
/* 操作结果：依次对L的每个数据元素输出 */
Status ListTraverse(LinkList L)
{
    LinkList p = L->next;
    while (p)
    {
        visit(p->data);
        printf("\t");
        p=p->next;
    }
    printf("\n");
 
    return OK;
}
 
/* 初始条件：单链表L已存在 */
/* 操作结果：若L为空表，则返回TRUE,否则返回FALSE */
Status ListEmpty(LinkList L)
{
    if (L->next)
    {
        return FALSE;
    }
    else
    {
        return TRUE;
    }
}
 
 
/* 初始条件：单链表L已存在 */
/* 操作结果：将L重置为空表 */
Status ClearList(LinkList *L)
{
    LinkList p, q;
    p = (*L)->next;      /* p指向第一个结点 */
    while (p)           /* 没到表尾 */
    {
        q = p->next;
        free(p);
        p = q;
    }
    (*L)->next = NULL;   /* 头结点指针域置为空 */
 
    return OK;
}
 
/* 初始条件：单链表已存在，且1 <= i <= ListLength */
/* 操作结果：用e返回第i个数据元素的值 */
Status GetElem(LinkList L,int i, ElemType *e)
{
    int j;
    LinkList p;     /* 声明一结点p */
    p = L->next; /* 让p指向链表L的第一个结点 */
    j = 1;          /* j为计数器 */
    while (p && j < i)       /* p不为空，或者计数器j还没有等于i时，循环继续 */
    {
        p = p->next;     /* 让p指向下一个结点 */
        j++;
    }
    if (!p || j > i)     /* 第i个元素不存在 */
    {
        return ERROR;
    }
 
    *e = p->data;            /* 取第i个元素的数据 */
    return OK;
}
 
 
/* 初始条件：单链表L已存在 */
/* 操作结果：返回L中第1个与与e满足关系的元素的下标；若没找到，则返回0 */
int LocateElem(LinkList L, ElemType e)
{
    int i = 0;
    LinkList p = L->next;
    while (p)
    {
        i++;
        if (p->data == e)    /* 找到了该数据元素 */
        {
            return i;
        }
        p = p->next;
    }
 
    return 0;
}
 
/* 初始条件：单链表已存在，且1 <= i <= ListLength */
/* 操作结果：删除L中的第i个数据元素，并且用e返回其值，L的长度减1 */
Status ListDelete(LinkList *L, int i, ElemType *e)
{
    int j;
    LinkList p, q;
    p = *L;
    j = 1;
    while (p->next && j < i)      /* 遍历寻找第i个元素 */
    {
        p = p->next;
        j++;
    }
    if (!(p->next) || j > i)
    {
        return ERROR;           /* 第i个元素不存在 */
    }
    q = p->next;
    p->next = q->next;            /* 标准删除语句，把q的后继结点赋值给p的后继结点 */
    *e = q->data;                /* 将q结点中的数据返回给e */
    free(q);
 
    return OK;
}
 
 
/* 头插法：随机产生n个元素的值，建立带头结点的单链线性表. */
void CreateListHead(LinkList *L, int n)
{
    LinkList p;
    int i;
    srand(time(0));     /* 初始化随机数种子 */
    *L = (LinkList)malloc(sizeof(Node));
    (*L)->next=NULL; /* 首先创建一个带头结点的单链表 */
 
    for (i = 0; i < n; i++)
    {
        p = (LinkList)malloc(sizeof(Node));     /* 生成新结点 */
        p->data=rand()%100 + 1;          /* 随机生成100以内的数字 */
        p->next = (*L)->next;
        (*L)->next = p;              /* 插入到表头 */
    }
}
 
 
/* 尾插法：随机产生n个元素的值，建立带表头结点的单线性链表. */
void CreateListTail(LinkList *L, int n)
{
    LinkList p, r;
    int i;
    srand(time(0));                         /* 初始化随机数种子 */
    *L = (LinkList)malloc(sizeof(Node));    /* L为整个线性表 */
    r = *L;             /* r为指向尾部的结点 */
    for (i = 0; i < n; i++)
    {
        p = (LinkList)malloc(sizeof(Node));         /* 生成新结点 */
        p->data=rand() % 100 + 1;                    /* 随机生成100以内的数字 */
        r->next = p;     /* 将表尾终端结点的指针指向新结点 */
        r = p;              /* 将当前新结点定义为表尾终端结点 */
    }
 
    r->next = NULL;          /* 表示当前链表结束 */
}
 
/* 测试方法，主函数 */
int main()
{
    LinkList L;
    ElemType e;
    Status i;
    int j, k;
    i = InitList(&L);
     
    printf("初始化L后：ListLength = %d\n",ListLength(L));
 
    for (j = 1; j <= 5; j++)
    {
        i = ListInsert(&L,1,j);
    }
    printf("在L的表头依次插入1~5后：L->data = ");
    ListTraverse(L);
 
    printf("ListLength(L) = %d \n",ListLength(L));
 
    i = ListEmpty(L);
    printf("L是否为空 (1:是 0:否)：i = %d \n",i);
 
    i = ClearList(&L);
    printf("清空L后：ListLength(L) = %d \n",ListLength(L));
    i = ListEmpty(L);
    printf("L是否为空 (1:是 0:否)：i = %d \n",i);
 
 
    for (j = 1; j <= 10; j++)
    {
        i = ListInsert(&L,j,j);
    }
    printf("在L的表头依次插入1~5后：L->data = ");
    ListTraverse(L);
 
    printf("ListLength(L) = %d \n",ListLength(L));
 
    ListInsert(&L,1,0);
    printf("在L的表头插入0后：L->data = ");
    ListTraverse(L);
    printf("ListLength(L) = %d \n",ListLength(L));
 
    GetElem(L,5,&e);
    printf("第5个元素的值为：%d \n",e);
 
    for (j = 3; j <= 4; j++)
    {
        k = LocateElem(L,j);
        if (k)
        {
            printf("第%d个元素的值为：%d \n",k,j);
        }
        else
        {
            printf("没有值为%d的元素. \n",j);
        }
    }
 
    printf("ListLength(L) = %d \n",ListLength(L));
 
    k = ListLength(L);      /* K为表长 = 11 */
    for (j = k + 1; j >= k; j--)
    {
        i = ListDelete(&L,j,&e);    /* 删除第j个数据元素 */
        if (i == 0)
        {
            printf("删除第%d个数据元素失败. \n");
        }
        else
        {
            printf("删除第%d个元素的值为：%d \n",j,e);
        }
    }
    printf("依次输出L中的元素：");
    ListTraverse(L);
    printf("ListLength(L) = %d \n",ListLength(L));
 
    j = 5;
    ListDelete(&L,j,&e);        /* 删除第5个数据元素 */
    printf("删除第%d个元素的值为：%d \n",j,e);
    printf("依次输出L中的元素：");
    ListTraverse(L);
    printf("ListLength(L) = %d \n",ListLength(L));
 
    i = ClearList(&L);
    printf("L是否为空 (1:是 0:否)：i = %d \n",i);
    printf("清空L后：ListLength(L) = %d \n",ListLength(L));
    /*  清空了.  */
     
    CreateListHead(&L,20);
    printf("整体创建L的元素(头插法): ");
    ListTraverse(L);
 
    i = ClearList(&L);
    printf("L是否为空 (1:是 0:否)：i = %d \n",i);
    printf("清空L后：ListLength(L) = %d \n",ListLength(L));
    /*  清空了.  */
 
    CreateListTail(&L,20);
    printf("整体创建L的元素(尾插法): ");
    ListTraverse(L);
 
    system("pause");
    return 0;
}

