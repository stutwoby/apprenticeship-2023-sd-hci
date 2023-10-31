namespace LitvaKebabs.Models
{
    public class PostcodeModel
    {
        public class Codes
        {
            public string admin_district { get; set; } = string.Empty;
            public string admin_county { get; set; } = string.Empty;
            public string admin_ward { get; set; } = string.Empty;
            public string parish { get; set; } = string.Empty;
            public string parliamentary_constituency { get; set; } = string.Empty;
            public string ccg { get; set; } = string.Empty;
            public string ccg_id { get; set; } = string.Empty;
            public string ced { get; set; } = string.Empty;
            public string nuts { get; set; } = string.Empty;
            public string lsoa { get; set; } = string.Empty;
            public string msoa { get; set; } = string.Empty;
            public string lau2 { get; set; } = string.Empty;
            public string pfa { get; set; } = string.Empty;
        }

        public class Result
        {
            public string postcode { get; set; } = string.Empty;
            public int? quality { get; set; }
            public int? eastings { get; set; }
            public int? northings { get; set; }
            public string country { get; set; } = string.Empty;
            public string nhs_ha { get; set; } = string.Empty;
            public double longitude { get; set; }
            public double latitude { get; set; }
            public string european_electoral_region { get; set; } = string.Empty;
            public string primary_care_trust { get; set; } = string.Empty;
            public string region { get; set; } = string.Empty;
            public string lsoa { get; set; } = string.Empty;
            public string msoa { get; set; } = string.Empty;
            public string incode { get; set; } = string.Empty;
            public string outcode { get; set; } = string.Empty;
            public string parliamentary_constituency { get; set; } = string.Empty;
            public string admin_district { get; set; } = string.Empty;
            public string parish { get; set; } = string.Empty;
            public string admin_county { get; set; } = string.Empty;
            public string date_of_introduction { get; set; } = string.Empty;
            public string admin_ward { get; set; } = string.Empty;
            public string ced { get; set; } = string.Empty;
            public string ccg { get; set; } = string.Empty;
            public string nuts { get; set; } = string.Empty;
            public string pfa { get; set; } = string.Empty;
            public Codes codes { get; set; }
        }

        public class Root
        {
            public int? status { get; set; }
            public Result result { get; set; }
        }


    }
}
