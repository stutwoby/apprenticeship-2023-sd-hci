package TakeAwayApp.KebabCreator;
import TakeAwayApp.*;
import TakeAwayApp.MenuHandler.MenuItems;
import TakeAwayApp.MenuHandler.MenuPrices;
import TakeAwayApp.ReceiptHandler.Total;

public class AddSaladKebab {
    DialogHandler dialog = new DialogHandler();

    MenuItems menuItems = new MenuItems();

    MenuPrices menuPrices = new MenuPrices();


    AddSaladKebab(String dialog){
        addSalad(dialog);
    }

    private void addSalad(String dialogHandler) {
        if (dialogHandler.equals("yes")) {
            Total.setItem(this.menuItems.getAnItem(4,1));
            Total.setPrice(this.menuPrices.getPrice(4,1));
            new AddSauceKebab();
        }
        else if (dialogHandler.equals("no"))
            new AddSauceKebab();
        else {
            System.out.println("Warning: Not an option");
            addSalad(dialog.getDialog("Would you like salad? (free)"));
        }
    }
}