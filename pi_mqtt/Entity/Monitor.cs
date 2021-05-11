namespace pi_mqtt.Entity
{
    public class Monitor
    {
        public int Id { get; set; } = 0;
        public string MemTotal { get; set; }
        public string MemFree { get; set; }
        public string MemAvailable { get; set; }
        public string Type { get; set; }
        public string Total { get; set; }
        public string Used { get; set; }
        public string Free { get; set; }
        public string UsedPercent { get; set; }
        public string Uptime { get; set; }
        public string Payload1 { get; set; }
        public string Payload5 { get; set; }
        public string Payload15 { get; set; }
    }
}