
#include <iostream>
#include <string>
using namespace std;
 
int f[2001][2001];
int n, m;
string a, b;
 
int main() {
	cin >> a >> b;
	n = a.size();
	m = b.size();
	f[0][0] = 0;
	for(int i=1; i<=n; i++)
		for(int j=1; j<=m; j++)
			if(a[i-1] == a[j-1]) f[i][j] = f[i-1][j-1] + 1;
			else f[i][j] = max(f[i-1][j], f[i][j-1]);
	cout << f[n][m] << endl;
	return 0;
}