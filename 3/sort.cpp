#include <iostream>
using namespace std;

class Sort {
    private:
       int a[100];
       int current_length ;
       int sorted_arr[100];
    public:
        void set_arr(int temp[],int len);
        void show_arr();
        void bubble_sort();
        void select_sort();
        void insert_sort_stupid();
        void insert_sort();
        void shell_sort();
        //int * bubble_sort(int arr[],int len);
};

// 设置私有待排序数组
void Sort::set_arr(int temp[],int len){
    int i = 0;
    for( i = 0; i < len; i ++) {
        this->a[i] = temp[i];
    }
    this->current_length = len;
}

// 显示待排序数组
void Sort::show_arr(){
    int i;
    cout << "show arr :" << endl;
    for(i = 0 ; i < this->current_length; i ++) {
        cout <<  this->a[i] << "-";
    }
    cout << endl;

}

// 冒泡排序
void  Sort::bubble_sort() {
    cout << "the length of the array which is waiting for sort:" <<endl;
    cout << this->current_length << endl;
    this->show_arr();
    int i,j;
    for(i=0; i<this->current_length-1; i++) {
        for(j = 0; j < this->current_length-i-1;j++) {
            if(this->a[j] > this->a[j+1]) {
                int temp;
                temp = this->a[j];
                this->a[j] = this->a[j+1];
                this->a[j+1] = temp;
            }
        }
    }
    cout << "after bubble sort:" << endl;
    this->show_arr();
}

// 选择排序
void  Sort::select_sort(){
    cout << "the length of the array which is waiting for select sort:" <<endl;
    cout << this->current_length << endl;
    this->show_arr();
    int i,j;
    for(i=0; i < this->current_length-1; i++) {
        int min_index = i;
        int min = this->a[i];
        for(j = i+1; j < this->current_length; j++) {
            if(this->a[j] < min) {
                min = this->a[j];
                min_index = j;
            }
        }
        if(min_index != i) {
            int temp;
            temp = this->a[i];
            this->a[i] = this->a[min_index];
            this->a[min_index] = temp;
        }
    }
    cout << "after select sort:" << endl;
    this->show_arr();
}

//插入排序 这里用的方式太笨，可以直接在原数组依次从a[1]开始插入之前的数组，这样就不用引入额外的临时数据，减少了空间消耗
void Sort::insert_sort_stupid(){
    int len = this->current_length;
    int temp[100] ;
    int sorted_count = 0;

    int i,j;
    for( i = 0; i < len; i++) {
        if(i == 0) {
            temp[i] = this->a[i];
            continue;
        }
        for(j = i -1; j>=0; j--) {
            if(this->a[i] < temp[j]) {
                temp[j+1] = temp[j];
                continue;
            } else {
               break;
            }
        }
        temp[j+1] = this->a[i];
    }
    cout << "insert sort:"<<endl;
    for( i = 0; i < len; i++) {
        this->a[i] = temp[i];
    }
}

//不使用额外的数组空间的插入排序
void Sort::insert_sort(){
    int len = this->current_length;

    int i,j;
    for( i = 1; i < len; i++) {
        int temp = this->a[i]; //注意在插入过程中，a[i]的值可能被前面的数字替代
        for(j = i -1; j>=0; j--) {
            if(temp < this->a[j]) {  
                this->a[j+1] = this->a[j]; // 此时a[i]处的值发生变化，故而需要使用temp进行比较
                continue;
            } else {
               break;
            }
        }
        this->a[j+1] = temp;
    }

    cout << "after insert sort:" << endl;
    this->show_arr();
}

//希尔排序
void Sort::shell_sort(){
    int gap;
    int len = this->current_length;
    for(gap = len/2; gap >0; gap/=2 ) {
        int i,j;
        for( i = gap; i < len; i ++) {
            int temp = this->a[i];
            for(j=i-gap; j>=0&& this->a[j] > temp; j-=gap) {
                this->a[j+gap] = this->a[j];
            }
            a[j+gap] = temp;
        }
    }
    this->show_arr();
}

int main(int argc, char* argv[]){
    
    Sort bubble,select;
    int temp[] = {1,2,4,5,3,9,7,2,1};
    // bubble.set_arr(temp,9);
    // bubble.bubble_sort();
    // select.set_arr(temp,9);
    // select.select_sort();

    Sort insert;
    insert.set_arr(temp,9);
    //insert.insert_sort();
    insert.shell_sort();

    
    return 0;
}