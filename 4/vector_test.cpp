#include<iostream>
#include<vector>
using namespace std;
int main(int argc,char* argv[]) {
    vector<int> v1,v2(10);
    for(int i = 1; i < 10; i++) {
        v1.push_back(i);
    }

    vector<int>::iterator it;
    for(it = v1.begin(); it <= v1.end(); it++) {
        cout << *it << endl;
    }

    // capacity 获取容器容量
    // size获取容器当前已使用容量
    // clear 清空容器内容
    // resize 改变当前已使用容量大小，如果是缩小，则超过容量部分数据会丢失
    // reserve 改变容器容量大小
    //push_back 放入新数据
    // pop_back 弹出最后一个数据

    cout << "v2.capacity:" << v2.capacity()<<endl; 
    cout << "v2.size:" << v2.size() <<endl;
    v2.clear();
 cout << " after clear v2.capacity:" << v2.capacity()<<endl;
    cout << "after clear v2.size:" << v2.size() <<endl;
    v2.resize(5);
     cout << " after resize(5) v2.capacity:" << v2.capacity()<<endl;
    cout << "after resize(5) v2.size:" << v2.size() <<endl;
    v2.reserve(16);
    v2.pop_back();
         cout << " after reserve(9) v2.capacity:" << v2.capacity()<<endl;
    cout << "after reserve(9) v2.size:" << v2.size() <<endl;
    v2.insert(v2.begin(),5);
    v2.insert(v2.end(),9);
    for (int i = 0; i <v2.size(); i++) {
        cout <<"v2[" << i << "]:" << v2[i] << endl;
    }

    
    
    return 0;
}