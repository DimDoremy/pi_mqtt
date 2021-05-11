namespace pi_mqtt.Entity
{
    public class Response
    {
        public Response()
        {
            ResData = null;
            ResCode = 0;
        }

        public Response(object resData, int resCode)
        {
            ResData = resData;
            ResCode = resCode;
        }

        public object ResData { get; set; }
        public int ResCode { get; set; }
    }
}