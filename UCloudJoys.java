import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;

public class UCloudJoys {

	public static void main(String[] args) throws IOException {
		BufferedReader br = new BufferedReader(new FileReader("UCloud.txt"));
		int total = 0;
		int i = 0;
		String text = null;
		while ((text = br.readLine()) != null) {
			int count = text.split("UCanUup").length - 1;
			System.out.println("Line " + i + " : " + count);
			total += count;
			i++;
		}
		System.out.println("Total: " + total);
		br.close();
	}
}