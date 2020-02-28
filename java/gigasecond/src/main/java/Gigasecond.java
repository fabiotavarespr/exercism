import java.time.*;

class Gigasecond {

    private LocalDateTime localDateTime;
    private static final long GIGASECOND = (long) Math.pow(10,9);;

    Gigasecond(LocalDateTime moment){
        this.localDateTime = moment;
    }

    Gigasecond(LocalDate moment) {
        this(moment.atStartOfDay());
    }

    LocalDateTime getDateTime() {
            return localDateTime.plusSeconds(GIGASECOND);
    }
}