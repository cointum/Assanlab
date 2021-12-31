import java.util.*;

public class Main {


	public class JSON{

		public int count;
		public String domainName;


	}
    
	public static List<String> subdomainVisits(String[] aDomains) {
		HashMap<String, Integer> map = new HashMap<String, Integer>();
		for (String str : aDomains) {
			StringBuilder sb = new StringBuilder();
			int i = 0;
			while(str.charAt(i) != ' ') {
				sb.append(str.charAt(i));
				i++;
			}
			int num = Integer.parseInt(sb.toString());
			int j = str.length() - 1;
			String s = "";
			while(j >= i) {
				char ch = str.charAt(j);
				if (ch == '.' || ch == ' ') {
					if(map.containsKey(s)) {
						map.put(s, map.get(s) + num);
					}else {
						map.put(s, num);
					}
				}
				s = ch + s;
				j--;
			}
		}
		List<String> ans = new ArrayList<String>();
		for(String str : map.keySet()) {
			String s = map.get(str) + " " + str;
			ans.add(s);
		}
		return ans;
	}

	public static void main(String[] args) {

		// taking input 
		
    String [] cpdomains  = new String[]{ "900 google.com","60 mail.yahoo.com","10 mobile.sports.yahoo.com" ,
	"40 sports.yahoo.com","300 yahoo.com", "10 stackoverflow.com","20 overflow.com",
	"5 com.com","2 en.wikipedia.org","1 m.wikipedia.org","1 mobile.sports","1 google.co.uk"};
	// JSON dom[]=new JSON[5];

	// dom[0].count=900;
	// dom[0].domainName="google.com";
		
		List<String> ans = subdomainVisits(cpdomains);
		Collections.sort(ans);
		for(String str : ans){
			System.out.println(str);
		}
	}

}