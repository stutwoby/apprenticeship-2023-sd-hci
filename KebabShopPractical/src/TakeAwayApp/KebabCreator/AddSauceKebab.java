package TakeAwayApp.KebabCreator;

import TakeAwayApp.*;
import TakeAwayApp.MenuHandler.MenuItems;
import TakeAwayApp.MenuHandler.MenuLoop;
import TakeAwayApp.MenuHandler.MenuPrices;
import TakeAwayApp.ReceiptHandler.Total;

public class AddSauceKebab {
    DialogHandler dialog = new DialogHandler();

    MenuItems menuItems = new MenuItems();

    MenuPrices menuPrices = new MenuPrices();
    static final String non = "Non";

    AddSauceKebab(){
        addSauceKebab();
    }

   private void addSauceKebab() {
        System.out.println(non);
        for (int i = 0; i < menuItems.getItemsList(3).length; i++) {
            System.out.println(menuItems.getAnItem(3,i));
        }
        String sauce = dialog.getDialog("What sauce would you like? (free)").toLowerCase();
        switch (sauce) {
            case"chilli": case"garlic mayo": case "mayo": case"ketchup": case"bbq":
                Total.setItem("      -" + sauce);
                Total.setPrice(menuPrices.getPrice(4,1));
                new MenuLoop(dialog.getDialog("Would you like anything else? yes/no"));
                break;
            case "non":
                new MenuLoop(dialog.getDialog("Would you like anything else? yes/no"));
                break;
            default:
                System.out.println("Warning: Not an option");
                addSauceKebab();
                break;
        }
    }
}