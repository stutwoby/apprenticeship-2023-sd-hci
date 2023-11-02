package TakeAwayApp.ReceiptHandler;

import TakeAwayApp.MenuHandler.MenuItems;
import TakeAwayApp.MenuHandler.MenuPrices;

public class ReceiptCreator {
    static MenuItems items = new MenuItems();
    static MenuPrices prices = new MenuPrices();


    ReceiptCreator() {}


    public static void addItemReceipt(String string) {
        for (int i = 0; i < items.getAllItemLists().length; i++)
            for (int j = 0; j < items.getItemsList(i).length; j++) {
                if (string.equalsIgnoreCase(items.getAnItem(i,j))) {
                    Total.setDelTime();
                    Total.setItem(items.getAnItem(i,j));
                    Total.setPrice(prices.getPrice(i,j));
                }
            }
    }
}