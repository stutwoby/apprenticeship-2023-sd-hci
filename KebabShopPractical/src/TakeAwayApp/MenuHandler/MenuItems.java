package TakeAwayApp.MenuHandler;

import TakeAwayApp.ReceiptHandler.ReceiptCreator;

public class MenuItems {

    public MenuItems() {
    }

    private final String[][] items =
            {
                    {"Donner Kebab", "Chicken Kebab", "Shish Kebab", "Seekh Kebab"},
                    {"Coke", "Pepsi", "Fanta", "7UP"},
                    {"Chips", "Cheese and Chips", "Onion Rings", "Curley Fries", "Chicken Nuggets x6", "Salad"},
                    {"Chilli", "Garlic Mayo", "Mayo", "Ketchup", "BBQ"},
                    {"      -large", "      -salad"}
            };

    public String[][] getAllItemLists(){return this.items;}

    public String[] getItemsList(int i) {return this.items[i];}

    public String getAnItem(int i, int j) {return this.items[i][j];}

    public boolean searchMenuItems(String menuItem) {
        for (int i = 0; i < this.items.length; i++)
            for (int j = 0; j < this.items[i].length; j++)
                if (menuItem.equals(this.items[i][j].toLowerCase())) {
                    ReceiptCreator.addItemReceipt(menuItem);
                    return true;
                }
        return false;
    }
}