public class Drink 
{
       //fields

    //constructors
    Drink(){
        System.out.println(" Please choose between: Coke, Lemonade, Orange juice or Iced Tea");
    }

    Drink(String drink){
        switch (drink){
            case "coke":
            Order.breakDown("coke", 1.00d);
            break;
            case "lemonade":
            Order.breakDown("lemonade",1.00d);
            break;
            case "orange juice":
            Order.breakDown("orange juice", 1.00d);
            break;
            case "iced tea":
            Order.breakDown("iced tea", 1.20d);
            break;

        }
    }
    
}
