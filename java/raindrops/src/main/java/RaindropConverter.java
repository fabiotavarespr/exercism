class RaindropConverter {

    String convert(int number) {
		
		String rainSound = new String();

		if (number%3 == 0) {
			rainSound += "Pling";
		}

		if (number%5 == 0) {
			rainSound += "Plang";
		}

		if (number%7 == 0) {
			rainSound += "Plong";
		}
		
		if (rainSound.isEmpty()) {
			rainSound = String.valueOf(number);
		}
		
		return rainSound;
    }
}