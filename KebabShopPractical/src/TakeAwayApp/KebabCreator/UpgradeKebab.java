package TakeAwayApp.KebabCreator;
import TakeAwayApp.*;
import TakeAwayApp.MenuHandler.MenuItems;
import TakeAwayApp.MenuHandler.MenuPrices;
import TakeAwayApp.ReceiptHandler.Total;


public class UpgradeKebab {

    MenuItems menuItems = new MenuItems();

    MenuPrices menuPrices = new MenuPrices();

    DialogHandler dialog = new DialogHandler();


  public UpgradeKebab(String kebab, String dialogHandler){
        largeKebab(kebab,dialogHandler);

    }


    private void largeKebab(String kebab, String dialogHandler) {
        if (dialogHandler.equals("yes")) {
            Total.setItem(menuItems.getAnItem(4, 0));
            Total.setPrice(menuPrices.getPrice(4, 0));
            new AddSaladKebab(dialog.getDialog("Would you like salad? (free)"));
        } else if (dialogHandler.equals("no"))
            new AddSaladKebab(dialog.getDialog("Would you like salad? (free)"));
        else {
            System.out.println("Warning: Not an option");
            largeKebab(kebab, dialog.getDialog("Would you like to go large? (£1.00)"));
        }
    }
}