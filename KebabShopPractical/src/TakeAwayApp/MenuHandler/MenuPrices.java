package TakeAwayApp.MenuHandler;

public class MenuPrices {

    private final float[][] prices =
            {
                    {6.00f, 6.90f, 6.80f, 7.00f},
                    {1.00f, 1.00f, 1.00f, 1.00f},
                    {1.60f, 2.10f, 2.00f, 1.60f, 3.00f, 1.20f},
                    {1.00f, 1.20f, 1.00f, 1.00f, 1.20f},
                    {1.00f, 0.00f}

            };


    public MenuPrices(){}


    public float[] getPricesList(int i){return this.prices[i];}

    public float getPrice(int i, int j){
        return this.prices[i][j];
    }


}
