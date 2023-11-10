package TakeAwayApp.MenuHandler;

public enum MenuBuilder {
    KEBABS {
        @Override
        public void list() {
            printItemsPrices(0);
        }
    },
    DRINKS {
        @Override
        public void list() {
            printItemsPrices(1);
        }
    },
    SIDES {
        @Override
        public void list() {
            printItemsPrices(2);
        }
    },
    SAUCES {
        @Override
        public void list() {
            printItemsPrices(3);
        }
    };

    public abstract void list();

    public void printItemsPrices(int i) {
        for (int j = 0; j < new MenuItems().getItemsList(i).length; j++) {
            System.out.printf(new MenuItems().getAnItem(i,j) + "  £" + "%.2f", new MenuPrices().getPrice(i,j));
            System.out.println("");
        }

    }
}