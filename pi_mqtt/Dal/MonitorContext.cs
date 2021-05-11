using System;
using Microsoft.EntityFrameworkCore;
using pi_mqtt.Entity;

namespace pi_mqtt.Dal
{
    public class MonitorContext : DbContext
    {
        public DbSet<Monitor> Monitor { get; set; }
        public DbSet<PC> PC { get; set; }

        protected override void OnConfiguring(DbContextOptionsBuilder optionsBuilder)
        {
            optionsBuilder.UseMySql(
                "Server=47.104.253.11;port=3306;Database=mqtt;uid=mqtt;pwd=186536_Wlj;Character Set=utf8;",new MySqlServerVersion(new Version(8, 0, 21)));
        }
    }
}