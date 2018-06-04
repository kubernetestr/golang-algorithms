public class Main {

    public static void main(String[] args) {
        int[] anArray;

        // allocates memory for 10 integers
        anArray = new int[101];

        int b;
        b = (int) (Math.random() * 100);
        System.out.println("random ="+b);

        int i = 0;
        while (i < b) {
            anArray[i] = i;
            //System.out.println(anArray[i]);
            i++;
        }
        while (b < 101) {
            anArray[b] = b - 1;
            //System.out.println(anArray[b]);
            b++;

        }
        int s=0;
        while(s<101){
            System.out.println(anArray[s]);
            s++;
        }
        int a=0;
        int c=0;
        while (a<101){
            if (anArray[a]==a) {
               // System.out.println("sadasdasd " + a);
            }

            else
                break;
                c=a+1;


                a++;

            }
        System.out.println("duplicate is "+ c);
        int k=0;
        while(k<101){
            if(anArray[k]==anArray[k+1]){

                break;
            }

            k++;




        }

        }




    }








    

