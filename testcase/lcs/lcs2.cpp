
#include<iostream>
using namespace std;
int CommonOrder(char x[] ,int m ,char y[], int n, char z[])
{
    int i,j,k;
    int L[10][10],S[10][10]; 
    for(j=0;j<=n;j++)
        L[0][j] = 0;
    for(i=0;i<=m;i++)
        L[i][0] = 0;
    for(i=1;i<=m;i++)
        for(j=1;j<=n;j++)
        {
            if(x[i-1]==y[j-1])
            {
                L[i][j] = L[i-1][j-1]+1;    //如果x[i-1]==y[j-1]，那么L[i][j]的值为其最近的左对角线元素的值加一 
                S[i][j] = 1;
            }
            else if(L[i][j-1]>=L[i-1][j])   //如果 L[i][j-1]>=L[i-1][j]即 L[i][j]左边相邻的元素大于等于上方相邻的                                            
            {                                //元素，则该元素的值取左方相邻元素的值
                L[i][j] = L[i][j-1];
                S[i][j] = 2;
            }    
            else                            //如果x[i-1]==y[j-1]且L[i][j-1]<L[i-1][j],则 L[i][j]的值为其上方的值 
            {
                L[i][j] = L[i-1][j];
                S[i][j] =3;
            }
        }
    for(j=0;j<=m;j++)
        {
        for(i=0;i<=n;i++)
            cout<<L[j][i];
        cout<<endl;
    }
    i=m;j=n;k=L[m][n];
    while(i>0 && j>0)
    {
        if(S[i][j] == 1)
        {
            z[k] = x[i];
            k--;
            i--;
            j--;
        }
        else if(S[i][j] == 2)
            j--;
        else 
            i--;
    }
    for(k=0;k<L[m][n];k++)
        cout<<z[k];
    cout<<endl;
    return L[m][n];    
}


int main()
{
    char x[]={'a','b','c','b','d','b'};
    char y[]={'a','c','b','b','a','b','d','b','b'};
    char b[50];
    cout<<"最长子序列的长度为："<<CommonOrder(x,6,y,9,b);
}