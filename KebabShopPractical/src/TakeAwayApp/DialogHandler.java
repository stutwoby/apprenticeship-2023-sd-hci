package TakeAwayApp;

import java.util.Scanner;

public class DialogHandler {
    private final Scanner scanner = new Scanner(System.in);


    public DialogHandler(){}


    private String dialog(String outPutMonologue) {
        System.out.println(outPutMonologue);
        String input = scanner.nextLine();
        return input.toLowerCase();
    }

    public String getDialog(String outPutMonologue){return dialog(outPutMonologue);}
}
