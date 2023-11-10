package TakeAwayApp.MenuHandler;

import TakeAwayApp.DialogHandler;
import TakeAwayApp.KebabCreator.UpgradeKebab;

public class ItemChecker {
    MenuItems menuItems = new MenuItems();
    DialogHandler dialogBox = new DialogHandler();

    public ItemChecker(String item){
       itemChecker(item);
    }

    private void itemChecker(String item){
        if(menuItems.searchMenuItems(item)) {
            switch (item) {
                case "donner kebab": case "chicken kebab": case "shish kebab": case "seekh kebab":
                new UpgradeKebab(item, dialogBox.getDialog("Would you like to go large for an extra £1.00"));
                break;
                default:
                new MenuLoop(dialogBox.getDialog("Would you like anything else? yes/no"));
                break;
            }
        }
        else{
            System.out.println("Warning: Item not on menu");
            new InitiateMenu();
        }
    }
}