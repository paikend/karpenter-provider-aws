//go:build !ignore_autogenerated

/*
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package pricing

// generated at 2024-12-16T13:14:38Z for us-east-1

import ec2types "github.com/aws/aws-sdk-go-v2/service/ec2/types"

var InitialOnDemandPricesAWS = map[string]map[ec2types.InstanceType]float64{
	// us-east-1
	"us-east-1": {
		// a1 family
		"a1.2xlarge": 0.204000, "a1.4xlarge": 0.408000, "a1.large": 0.051000, "a1.medium": 0.025500,
		"a1.metal": 0.408000, "a1.xlarge": 0.102000,
		// c1 family
		"c1.medium": 0.130000, "c1.xlarge": 0.520000,
		// c3 family
		"c3.2xlarge": 0.420000, "c3.4xlarge": 0.840000, "c3.8xlarge": 1.680000, "c3.large": 0.105000,
		"c3.xlarge": 0.210000,
		// c4 family
		"c4.2xlarge": 0.398000, "c4.4xlarge": 0.796000, "c4.8xlarge": 1.591000, "c4.large": 0.100000,
		"c4.xlarge": 0.199000,
		// c5 family
		"c5.12xlarge": 2.040000, "c5.18xlarge": 3.060000, "c5.24xlarge": 4.080000, "c5.2xlarge": 0.340000,
		"c5.4xlarge": 0.680000, "c5.9xlarge": 1.530000, "c5.large": 0.085000, "c5.metal": 4.080000,
		"c5.xlarge": 0.170000,
		// c5a family
		"c5a.12xlarge": 1.848000, "c5a.16xlarge": 2.464000, "c5a.24xlarge": 3.696000, "c5a.2xlarge": 0.308000,
		"c5a.4xlarge": 0.616000, "c5a.8xlarge": 1.232000, "c5a.large": 0.077000, "c5a.xlarge": 0.154000,
		// c5ad family
		"c5ad.12xlarge": 2.064000, "c5ad.16xlarge": 2.752000, "c5ad.24xlarge": 4.128000, "c5ad.2xlarge": 0.344000,
		"c5ad.4xlarge": 0.688000, "c5ad.8xlarge": 1.376000, "c5ad.large": 0.086000, "c5ad.xlarge": 0.172000,
		// c5d family
		"c5d.12xlarge": 2.304000, "c5d.18xlarge": 3.456000, "c5d.24xlarge": 4.608000, "c5d.2xlarge": 0.384000,
		"c5d.4xlarge": 0.768000, "c5d.9xlarge": 1.728000, "c5d.large": 0.096000, "c5d.metal": 4.608000,
		"c5d.xlarge": 0.192000,
		// c5n family
		"c5n.18xlarge": 3.888000, "c5n.2xlarge": 0.432000, "c5n.4xlarge": 0.864000, "c5n.9xlarge": 1.944000,
		"c5n.large": 0.108000, "c5n.metal": 3.888000, "c5n.xlarge": 0.216000,
		// c6a family
		"c6a.12xlarge": 1.836000, "c6a.16xlarge": 2.448000, "c6a.24xlarge": 3.672000, "c6a.2xlarge": 0.306000,
		"c6a.32xlarge": 4.896000, "c6a.48xlarge": 7.344000, "c6a.4xlarge": 0.612000, "c6a.8xlarge": 1.224000,
		"c6a.large": 0.076500, "c6a.metal": 7.344000, "c6a.xlarge": 0.153000,
		// c6g family
		"c6g.12xlarge": 1.632000, "c6g.16xlarge": 2.176000, "c6g.2xlarge": 0.272000, "c6g.4xlarge": 0.544000,
		"c6g.8xlarge": 1.088000, "c6g.large": 0.068000, "c6g.medium": 0.034000, "c6g.metal": 2.306600,
		"c6g.xlarge": 0.136000,
		// c6gd family
		"c6gd.12xlarge": 1.843200, "c6gd.16xlarge": 2.457600, "c6gd.2xlarge": 0.307200, "c6gd.4xlarge": 0.614400,
		"c6gd.8xlarge": 1.228800, "c6gd.large": 0.076800, "c6gd.medium": 0.038400, "c6gd.metal": 2.605100,
		"c6gd.xlarge": 0.153600,
		// c6gn family
		"c6gn.12xlarge": 2.073600, "c6gn.16xlarge": 2.764800, "c6gn.2xlarge": 0.345600, "c6gn.4xlarge": 0.691200,
		"c6gn.8xlarge": 1.382400, "c6gn.large": 0.086400, "c6gn.medium": 0.043200, "c6gn.xlarge": 0.172800,
		// c6i family
		"c6i.12xlarge": 2.040000, "c6i.16xlarge": 2.720000, "c6i.24xlarge": 4.080000, "c6i.2xlarge": 0.340000,
		"c6i.32xlarge": 5.440000, "c6i.4xlarge": 0.680000, "c6i.8xlarge": 1.360000, "c6i.large": 0.085000,
		"c6i.metal": 5.440000, "c6i.xlarge": 0.170000,
		// c6id family
		"c6id.12xlarge": 2.419200, "c6id.16xlarge": 3.225600, "c6id.24xlarge": 4.838400, "c6id.2xlarge": 0.403200,
		"c6id.32xlarge": 6.451200, "c6id.4xlarge": 0.806400, "c6id.8xlarge": 1.612800, "c6id.large": 0.100800,
		"c6id.metal": 6.451200, "c6id.xlarge": 0.201600,
		// c6in family
		"c6in.12xlarge": 2.721600, "c6in.16xlarge": 3.628800, "c6in.24xlarge": 5.443200, "c6in.2xlarge": 0.453600,
		"c6in.32xlarge": 7.257600, "c6in.4xlarge": 0.907200, "c6in.8xlarge": 1.814400, "c6in.large": 0.113400,
		"c6in.metal": 7.257600, "c6in.xlarge": 0.226800,
		// c7a family
		"c7a.12xlarge": 2.463360, "c7a.16xlarge": 3.284480, "c7a.24xlarge": 4.926720, "c7a.2xlarge": 0.410560,
		"c7a.32xlarge": 6.568960, "c7a.48xlarge": 9.853440, "c7a.4xlarge": 0.821120, "c7a.8xlarge": 1.642240,
		"c7a.large": 0.102640, "c7a.medium": 0.051320, "c7a.metal-48xl": 9.853440, "c7a.xlarge": 0.205280,
		// c7g family
		"c7g.12xlarge": 1.740000, "c7g.16xlarge": 2.320000, "c7g.2xlarge": 0.290000, "c7g.4xlarge": 0.580000,
		"c7g.8xlarge": 1.160000, "c7g.large": 0.072500, "c7g.medium": 0.036300, "c7g.metal": 2.320000,
		"c7g.xlarge": 0.145000,
		// c7gd family
		"c7gd.12xlarge": 2.177300, "c7gd.16xlarge": 2.903000, "c7gd.2xlarge": 0.362900, "c7gd.4xlarge": 0.725800,
		"c7gd.8xlarge": 1.451500, "c7gd.large": 0.090700, "c7gd.medium": 0.045400, "c7gd.metal": 2.903000,
		"c7gd.xlarge": 0.181400,
		// c7gn family
		"c7gn.12xlarge": 2.995200, "c7gn.16xlarge": 3.993600, "c7gn.2xlarge": 0.499200, "c7gn.4xlarge": 0.998400,
		"c7gn.8xlarge": 1.996800, "c7gn.large": 0.124800, "c7gn.medium": 0.062400, "c7gn.metal": 3.993600,
		"c7gn.xlarge": 0.249600,
		// c7i-flex family
		"c7i-flex.2xlarge": 0.339150, "c7i-flex.4xlarge": 0.678300, "c7i-flex.8xlarge": 1.356600,
		"c7i-flex.large": 0.084790, "c7i-flex.xlarge": 0.169580,
		// c7i family
		"c7i.12xlarge": 2.142000, "c7i.16xlarge": 2.856000, "c7i.24xlarge": 4.284000, "c7i.2xlarge": 0.357000,
		"c7i.48xlarge": 8.568000, "c7i.4xlarge": 0.714000, "c7i.8xlarge": 1.428000, "c7i.large": 0.089250,
		"c7i.metal-24xl": 4.712400, "c7i.metal-48xl": 8.568000, "c7i.xlarge": 0.178500,
		// c8g family
		"c8g.12xlarge": 1.914240, "c8g.16xlarge": 2.552320, "c8g.24xlarge": 3.828480, "c8g.2xlarge": 0.319040,
		"c8g.48xlarge": 7.656960, "c8g.4xlarge": 0.638080, "c8g.8xlarge": 1.276160, "c8g.large": 0.079760,
		"c8g.medium": 0.039880, "c8g.metal-24xl": 4.211330, "c8g.metal-48xl": 7.656960, "c8g.xlarge": 0.159520,
		// cr1 family
		"cr1.8xlarge": 3.500000,
		// d2 family
		"d2.2xlarge": 1.380000, "d2.4xlarge": 2.760000, "d2.8xlarge": 5.520000, "d2.xlarge": 0.690000,
		// d3 family
		"d3.2xlarge": 0.999000, "d3.4xlarge": 1.998000, "d3.8xlarge": 3.995520, "d3.xlarge": 0.499000,
		// d3en family
		"d3en.12xlarge": 6.308640, "d3en.2xlarge": 1.051000, "d3en.4xlarge": 2.103000, "d3en.6xlarge": 3.154000,
		"d3en.8xlarge": 4.205760, "d3en.xlarge": 0.526000,
		// dl1 family
		"dl1.24xlarge": 13.109040,
		// f1 family
		"f1.16xlarge": 13.200000, "f1.2xlarge": 1.650000, "f1.4xlarge": 3.300000,
		// f2 family
		"f2.12xlarge": 3.960000, "f2.48xlarge": 15.840000,
		// g2 family
		"g2.2xlarge": 0.650000, "g2.8xlarge": 2.600000,
		// g3 family
		"g3.16xlarge": 4.560000, "g3.4xlarge": 1.140000, "g3.8xlarge": 2.280000,
		// g3s family
		"g3s.xlarge": 0.750000,
		// g4ad family
		"g4ad.16xlarge": 3.468000, "g4ad.2xlarge": 0.541170, "g4ad.4xlarge": 0.867000, "g4ad.8xlarge": 1.734000,
		"g4ad.xlarge": 0.378530,
		// g4dn family
		"g4dn.12xlarge": 3.912000, "g4dn.16xlarge": 4.352000, "g4dn.2xlarge": 0.752000, "g4dn.4xlarge": 1.204000,
		"g4dn.8xlarge": 2.176000, "g4dn.metal": 7.824000, "g4dn.xlarge": 0.526000,
		// g5 family
		"g5.12xlarge": 5.672000, "g5.16xlarge": 4.096000, "g5.24xlarge": 8.144000, "g5.2xlarge": 1.212000,
		"g5.48xlarge": 16.288000, "g5.4xlarge": 1.624000, "g5.8xlarge": 2.448000, "g5.xlarge": 1.006000,
		// g5g family
		"g5g.16xlarge": 2.744000, "g5g.2xlarge": 0.556000, "g5g.4xlarge": 0.828000, "g5g.8xlarge": 1.372000,
		"g5g.metal": 2.744000, "g5g.xlarge": 0.420000,
		// g6 family
		"g6.12xlarge": 4.601600, "g6.16xlarge": 3.396800, "g6.24xlarge": 6.675200, "g6.2xlarge": 0.977600,
		"g6.48xlarge": 13.350400, "g6.4xlarge": 1.323200, "g6.8xlarge": 2.014400, "g6.xlarge": 0.804800,
		// g6e family
		"g6e.12xlarge": 10.492640, "g6e.16xlarge": 7.577190, "g6e.24xlarge": 15.065590, "g6e.2xlarge": 2.242080,
		"g6e.48xlarge": 30.131180, "g6e.4xlarge": 3.004240, "g6e.8xlarge": 4.528560, "g6e.xlarge": 1.861000,
		// gr6 family
		"gr6.4xlarge": 1.539200, "gr6.8xlarge": 2.446400,
		// h1 family
		"h1.16xlarge": 3.744000, "h1.2xlarge": 0.468000, "h1.4xlarge": 0.936000, "h1.8xlarge": 1.872000,
		// hpc7g family
		"hpc7g.16xlarge": 1.683200, "hpc7g.4xlarge": 1.683200, "hpc7g.8xlarge": 1.683200,
		// i2 family
		"i2.2xlarge": 1.705000, "i2.4xlarge": 3.410000, "i2.8xlarge": 6.820000, "i2.xlarge": 0.853000,
		// i3 family
		"i3.16xlarge": 4.992000, "i3.2xlarge": 0.624000, "i3.4xlarge": 1.248000, "i3.8xlarge": 2.496000,
		"i3.large": 0.156000, "i3.metal": 4.992000, "i3.xlarge": 0.312000,
		// i3en family
		"i3en.12xlarge": 5.424000, "i3en.24xlarge": 10.848000, "i3en.2xlarge": 0.904000, "i3en.3xlarge": 1.356000,
		"i3en.6xlarge": 2.712000, "i3en.large": 0.226000, "i3en.metal": 10.848000, "i3en.xlarge": 0.452000,
		// i4g family
		"i4g.16xlarge": 4.942080, "i4g.2xlarge": 0.617760, "i4g.4xlarge": 1.235520, "i4g.8xlarge": 2.471040,
		"i4g.large": 0.154440, "i4g.xlarge": 0.308880,
		// i4i family
		"i4i.12xlarge": 4.118000, "i4i.16xlarge": 5.491000, "i4i.24xlarge": 8.236800, "i4i.2xlarge": 0.686000,
		"i4i.32xlarge": 10.982400, "i4i.4xlarge": 1.373000, "i4i.8xlarge": 2.746000, "i4i.large": 0.172000,
		"i4i.metal": 10.982000, "i4i.xlarge": 0.343000,
		// i7ie family
		"i7ie.12xlarge": 6.237600, "i7ie.18xlarge": 9.356400, "i7ie.24xlarge": 12.475200, "i7ie.2xlarge": 1.039600,
		"i7ie.3xlarge": 1.559400, "i7ie.48xlarge": 24.950400, "i7ie.6xlarge": 3.118800, "i7ie.large": 0.259900,
		"i7ie.xlarge": 0.519800,
		// i8g family
		"i8g.12xlarge": 4.118400, "i8g.16xlarge": 5.491200, "i8g.24xlarge": 8.236800, "i8g.2xlarge": 0.686400,
		"i8g.4xlarge": 1.372800, "i8g.8xlarge": 2.745600, "i8g.large": 0.171600, "i8g.metal-24xl": 9.060480,
		"i8g.xlarge": 0.343200,
		// im4gn family
		"im4gn.16xlarge": 5.820670, "im4gn.2xlarge": 0.727580, "im4gn.4xlarge": 1.455170, "im4gn.8xlarge": 2.910340,
		"im4gn.large": 0.181900, "im4gn.xlarge": 0.363790,
		// inf1 family
		"inf1.24xlarge": 4.721000, "inf1.2xlarge": 0.362000, "inf1.6xlarge": 1.180000, "inf1.xlarge": 0.228000,
		// inf2 family
		"inf2.24xlarge": 6.490630, "inf2.48xlarge": 12.981270, "inf2.8xlarge": 1.967860, "inf2.xlarge": 0.758200,
		// is4gen family
		"is4gen.2xlarge": 1.152600, "is4gen.4xlarge": 2.305200, "is4gen.8xlarge": 4.610400,
		"is4gen.large": 0.288150, "is4gen.medium": 0.144080, "is4gen.xlarge": 0.576300,
		// m1 family
		"m1.large": 0.175000, "m1.medium": 0.087000, "m1.small": 0.044000, "m1.xlarge": 0.350000,
		// m2 family
		"m2.2xlarge": 0.490000, "m2.4xlarge": 0.980000, "m2.xlarge": 0.245000,
		// m3 family
		"m3.2xlarge": 0.532000, "m3.large": 0.133000, "m3.medium": 0.067000, "m3.xlarge": 0.266000,
		// m4 family
		"m4.10xlarge": 2.000000, "m4.16xlarge": 3.200000, "m4.2xlarge": 0.400000, "m4.4xlarge": 0.800000,
		"m4.large": 0.100000, "m4.xlarge": 0.200000,
		// m5 family
		"m5.12xlarge": 2.304000, "m5.16xlarge": 3.072000, "m5.24xlarge": 4.608000, "m5.2xlarge": 0.384000,
		"m5.4xlarge": 0.768000, "m5.8xlarge": 1.536000, "m5.large": 0.096000, "m5.metal": 4.608000,
		"m5.xlarge": 0.192000,
		// m5a family
		"m5a.12xlarge": 2.064000, "m5a.16xlarge": 2.752000, "m5a.24xlarge": 4.128000, "m5a.2xlarge": 0.344000,
		"m5a.4xlarge": 0.688000, "m5a.8xlarge": 1.376000, "m5a.large": 0.086000, "m5a.xlarge": 0.172000,
		// m5ad family
		"m5ad.12xlarge": 2.472000, "m5ad.16xlarge": 3.296000, "m5ad.24xlarge": 4.944000, "m5ad.2xlarge": 0.412000,
		"m5ad.4xlarge": 0.824000, "m5ad.8xlarge": 1.648000, "m5ad.large": 0.103000, "m5ad.xlarge": 0.206000,
		// m5d family
		"m5d.12xlarge": 2.712000, "m5d.16xlarge": 3.616000, "m5d.24xlarge": 5.424000, "m5d.2xlarge": 0.452000,
		"m5d.4xlarge": 0.904000, "m5d.8xlarge": 1.808000, "m5d.large": 0.113000, "m5d.metal": 5.424000,
		"m5d.xlarge": 0.226000,
		// m5dn family
		"m5dn.12xlarge": 3.264000, "m5dn.16xlarge": 4.352000, "m5dn.24xlarge": 6.528000, "m5dn.2xlarge": 0.544000,
		"m5dn.4xlarge": 1.088000, "m5dn.8xlarge": 2.176000, "m5dn.large": 0.136000, "m5dn.metal": 6.528000,
		"m5dn.xlarge": 0.272000,
		// m5n family
		"m5n.12xlarge": 2.856000, "m5n.16xlarge": 3.808000, "m5n.24xlarge": 5.712000, "m5n.2xlarge": 0.476000,
		"m5n.4xlarge": 0.952000, "m5n.8xlarge": 1.904000, "m5n.large": 0.119000, "m5n.metal": 5.712000,
		"m5n.xlarge": 0.238000,
		// m5zn family
		"m5zn.12xlarge": 3.964100, "m5zn.2xlarge": 0.660700, "m5zn.3xlarge": 0.991000, "m5zn.6xlarge": 1.982000,
		"m5zn.large": 0.165200, "m5zn.metal": 4.360500, "m5zn.xlarge": 0.330300,
		// m6a family
		"m6a.12xlarge": 2.073600, "m6a.16xlarge": 2.764800, "m6a.24xlarge": 4.147200, "m6a.2xlarge": 0.345600,
		"m6a.32xlarge": 5.529600, "m6a.48xlarge": 8.294400, "m6a.4xlarge": 0.691200, "m6a.8xlarge": 1.382400,
		"m6a.large": 0.086400, "m6a.metal": 8.294400, "m6a.xlarge": 0.172800,
		// m6g family
		"m6g.12xlarge": 1.848000, "m6g.16xlarge": 2.464000, "m6g.2xlarge": 0.308000, "m6g.4xlarge": 0.616000,
		"m6g.8xlarge": 1.232000, "m6g.large": 0.077000, "m6g.medium": 0.038500, "m6g.metal": 2.611200,
		"m6g.xlarge": 0.154000,
		// m6gd family
		"m6gd.12xlarge": 2.169600, "m6gd.16xlarge": 2.892800, "m6gd.2xlarge": 0.361600, "m6gd.4xlarge": 0.723200,
		"m6gd.8xlarge": 1.446400, "m6gd.large": 0.090400, "m6gd.medium": 0.045200, "m6gd.metal": 3.066400,
		"m6gd.xlarge": 0.180800,
		// m6i family
		"m6i.12xlarge": 2.304000, "m6i.16xlarge": 3.072000, "m6i.24xlarge": 4.608000, "m6i.2xlarge": 0.384000,
		"m6i.32xlarge": 6.144000, "m6i.4xlarge": 0.768000, "m6i.8xlarge": 1.536000, "m6i.large": 0.096000,
		"m6i.metal": 6.144000, "m6i.xlarge": 0.192000,
		// m6id family
		"m6id.12xlarge": 2.847600, "m6id.16xlarge": 3.796800, "m6id.24xlarge": 5.695200, "m6id.2xlarge": 0.474600,
		"m6id.32xlarge": 7.593600, "m6id.4xlarge": 0.949200, "m6id.8xlarge": 1.898400, "m6id.large": 0.118650,
		"m6id.metal": 7.593600, "m6id.xlarge": 0.237300,
		// m6idn family
		"m6idn.12xlarge": 3.818880, "m6idn.16xlarge": 5.091840, "m6idn.24xlarge": 7.637760,
		"m6idn.2xlarge": 0.636480, "m6idn.32xlarge": 10.183680, "m6idn.4xlarge": 1.272960, "m6idn.8xlarge": 2.545920,
		"m6idn.large": 0.159120, "m6idn.metal": 10.183680, "m6idn.xlarge": 0.318240,
		// m6in family
		"m6in.12xlarge": 3.341520, "m6in.16xlarge": 4.455360, "m6in.24xlarge": 6.683040, "m6in.2xlarge": 0.556920,
		"m6in.32xlarge": 8.910720, "m6in.4xlarge": 1.113840, "m6in.8xlarge": 2.227680, "m6in.large": 0.139230,
		"m6in.metal": 8.910720, "m6in.xlarge": 0.278460,
		// m7a family
		"m7a.12xlarge": 2.782080, "m7a.16xlarge": 3.709440, "m7a.24xlarge": 5.564160, "m7a.2xlarge": 0.463680,
		"m7a.32xlarge": 7.418880, "m7a.48xlarge": 11.128320, "m7a.4xlarge": 0.927360, "m7a.8xlarge": 1.854720,
		"m7a.large": 0.115920, "m7a.medium": 0.057960, "m7a.metal-48xl": 11.128320, "m7a.xlarge": 0.231840,
		// m7g family
		"m7g.12xlarge": 1.958400, "m7g.16xlarge": 2.611200, "m7g.2xlarge": 0.326400, "m7g.4xlarge": 0.652800,
		"m7g.8xlarge": 1.305600, "m7g.large": 0.081600, "m7g.medium": 0.040800, "m7g.metal": 2.611200,
		"m7g.xlarge": 0.163200,
		// m7gd family
		"m7gd.12xlarge": 2.562800, "m7gd.16xlarge": 3.417100, "m7gd.2xlarge": 0.427100, "m7gd.4xlarge": 0.854300,
		"m7gd.8xlarge": 1.708600, "m7gd.large": 0.106800, "m7gd.medium": 0.053400, "m7gd.metal": 3.417100,
		"m7gd.xlarge": 0.213600,
		// m7i-flex family
		"m7i-flex.2xlarge": 0.383040, "m7i-flex.4xlarge": 0.766080, "m7i-flex.8xlarge": 1.532160,
		"m7i-flex.large": 0.095760, "m7i-flex.xlarge": 0.191520,
		// m7i family
		"m7i.12xlarge": 2.419200, "m7i.16xlarge": 3.225600, "m7i.24xlarge": 4.838400, "m7i.2xlarge": 0.403200,
		"m7i.48xlarge": 9.676800, "m7i.4xlarge": 0.806400, "m7i.8xlarge": 1.612800, "m7i.large": 0.100800,
		"m7i.metal-24xl": 5.322240, "m7i.metal-48xl": 9.676800, "m7i.xlarge": 0.201600,
		// m8g family
		"m8g.12xlarge": 2.154240, "m8g.16xlarge": 2.872320, "m8g.24xlarge": 4.308480, "m8g.2xlarge": 0.359040,
		"m8g.48xlarge": 8.616960, "m8g.4xlarge": 0.718080, "m8g.8xlarge": 1.436160, "m8g.large": 0.089760,
		"m8g.medium": 0.044880, "m8g.metal-24xl": 4.739330, "m8g.metal-48xl": 8.616960, "m8g.xlarge": 0.179520,
		// p2 family
		"p2.16xlarge": 14.400000, "p2.8xlarge": 7.200000, "p2.xlarge": 0.900000,
		// p3 family
		"p3.16xlarge": 24.480000, "p3.2xlarge": 3.060000, "p3.8xlarge": 12.240000,
		// p3dn family
		"p3dn.24xlarge": 31.212000,
		// p4d family
		"p4d.24xlarge": 32.772600,
		// p4de family
		"p4de.24xlarge": 40.965750,
		// p5 family
		"p5.48xlarge": 98.320000,
		// r3 family
		"r3.2xlarge": 0.665000, "r3.4xlarge": 1.330000, "r3.8xlarge": 2.660000, "r3.large": 0.166000,
		"r3.xlarge": 0.333000,
		// r4 family
		"r4.16xlarge": 4.256000, "r4.2xlarge": 0.532000, "r4.4xlarge": 1.064000, "r4.8xlarge": 2.128000,
		"r4.large": 0.133000, "r4.xlarge": 0.266000,
		// r5 family
		"r5.12xlarge": 3.024000, "r5.16xlarge": 4.032000, "r5.24xlarge": 6.048000, "r5.2xlarge": 0.504000,
		"r5.4xlarge": 1.008000, "r5.8xlarge": 2.016000, "r5.large": 0.126000, "r5.metal": 6.048000,
		"r5.xlarge": 0.252000,
		// r5a family
		"r5a.12xlarge": 2.712000, "r5a.16xlarge": 3.616000, "r5a.24xlarge": 5.424000, "r5a.2xlarge": 0.452000,
		"r5a.4xlarge": 0.904000, "r5a.8xlarge": 1.808000, "r5a.large": 0.113000, "r5a.xlarge": 0.226000,
		// r5ad family
		"r5ad.12xlarge": 3.144000, "r5ad.16xlarge": 4.192000, "r5ad.24xlarge": 6.288000, "r5ad.2xlarge": 0.524000,
		"r5ad.4xlarge": 1.048000, "r5ad.8xlarge": 2.096000, "r5ad.large": 0.131000, "r5ad.xlarge": 0.262000,
		// r5b family
		"r5b.12xlarge": 3.576000, "r5b.16xlarge": 4.768000, "r5b.24xlarge": 7.152000, "r5b.2xlarge": 0.596000,
		"r5b.4xlarge": 1.192000, "r5b.8xlarge": 2.384000, "r5b.large": 0.149000, "r5b.metal": 7.867200,
		"r5b.xlarge": 0.298000,
		// r5d family
		"r5d.12xlarge": 3.456000, "r5d.16xlarge": 4.608000, "r5d.24xlarge": 6.912000, "r5d.2xlarge": 0.576000,
		"r5d.4xlarge": 1.152000, "r5d.8xlarge": 2.304000, "r5d.large": 0.144000, "r5d.metal": 6.912000,
		"r5d.xlarge": 0.288000,
		// r5dn family
		"r5dn.12xlarge": 4.008000, "r5dn.16xlarge": 5.344000, "r5dn.24xlarge": 8.016000, "r5dn.2xlarge": 0.668000,
		"r5dn.4xlarge": 1.336000, "r5dn.8xlarge": 2.672000, "r5dn.large": 0.167000, "r5dn.metal": 8.016000,
		"r5dn.xlarge": 0.334000,
		// r5n family
		"r5n.12xlarge": 3.576000, "r5n.16xlarge": 4.768000, "r5n.24xlarge": 7.152000, "r5n.2xlarge": 0.596000,
		"r5n.4xlarge": 1.192000, "r5n.8xlarge": 2.384000, "r5n.large": 0.149000, "r5n.metal": 7.152000,
		"r5n.xlarge": 0.298000,
		// r6a family
		"r6a.12xlarge": 2.721600, "r6a.16xlarge": 3.628800, "r6a.24xlarge": 5.443200, "r6a.2xlarge": 0.453600,
		"r6a.32xlarge": 7.257600, "r6a.48xlarge": 10.886400, "r6a.4xlarge": 0.907200, "r6a.8xlarge": 1.814400,
		"r6a.large": 0.113400, "r6a.metal": 10.886400, "r6a.xlarge": 0.226800,
		// r6g family
		"r6g.12xlarge": 2.419200, "r6g.16xlarge": 3.225600, "r6g.2xlarge": 0.403200, "r6g.4xlarge": 0.806400,
		"r6g.8xlarge": 1.612800, "r6g.large": 0.100800, "r6g.medium": 0.050400, "r6g.metal": 3.419100,
		"r6g.xlarge": 0.201600,
		// r6gd family
		"r6gd.12xlarge": 2.764800, "r6gd.16xlarge": 3.686400, "r6gd.2xlarge": 0.460800, "r6gd.4xlarge": 0.921600,
		"r6gd.8xlarge": 1.843200, "r6gd.large": 0.115200, "r6gd.medium": 0.057600, "r6gd.metal": 3.907600,
		"r6gd.xlarge": 0.230400,
		// r6i family
		"r6i.12xlarge": 3.024000, "r6i.16xlarge": 4.032000, "r6i.24xlarge": 6.048000, "r6i.2xlarge": 0.504000,
		"r6i.32xlarge": 8.064000, "r6i.4xlarge": 1.008000, "r6i.8xlarge": 2.016000, "r6i.large": 0.126000,
		"r6i.metal": 8.064000, "r6i.xlarge": 0.252000,
		// r6id family
		"r6id.12xlarge": 3.628800, "r6id.16xlarge": 4.838400, "r6id.24xlarge": 7.257600, "r6id.2xlarge": 0.604800,
		"r6id.32xlarge": 9.676800, "r6id.4xlarge": 1.209600, "r6id.8xlarge": 2.419200, "r6id.large": 0.151200,
		"r6id.metal": 9.676800, "r6id.xlarge": 0.302400,
		// r6idn family
		"r6idn.12xlarge": 4.689360, "r6idn.16xlarge": 6.252480, "r6idn.24xlarge": 9.378720,
		"r6idn.2xlarge": 0.781560, "r6idn.32xlarge": 12.504960, "r6idn.4xlarge": 1.563120, "r6idn.8xlarge": 3.126240,
		"r6idn.large": 0.195390, "r6idn.metal": 12.504960, "r6idn.xlarge": 0.390780,
		// r6in family
		"r6in.12xlarge": 4.183920, "r6in.16xlarge": 5.578560, "r6in.24xlarge": 8.367840, "r6in.2xlarge": 0.697320,
		"r6in.32xlarge": 11.157120, "r6in.4xlarge": 1.394640, "r6in.8xlarge": 2.789280, "r6in.large": 0.174330,
		"r6in.metal": 11.157120, "r6in.xlarge": 0.348660,
		// r7a family
		"r7a.12xlarge": 3.651600, "r7a.16xlarge": 4.868800, "r7a.24xlarge": 7.303200, "r7a.2xlarge": 0.608600,
		"r7a.32xlarge": 9.737600, "r7a.48xlarge": 14.606400, "r7a.4xlarge": 1.217200, "r7a.8xlarge": 2.434400,
		"r7a.large": 0.152150, "r7a.medium": 0.076080, "r7a.metal-48xl": 14.606400, "r7a.xlarge": 0.304300,
		// r7g family
		"r7g.12xlarge": 2.570400, "r7g.16xlarge": 3.427200, "r7g.2xlarge": 0.428400, "r7g.4xlarge": 0.856800,
		"r7g.8xlarge": 1.713600, "r7g.large": 0.107100, "r7g.medium": 0.053600, "r7g.metal": 3.427200,
		"r7g.xlarge": 0.214200,
		// r7gd family
		"r7gd.12xlarge": 3.265900, "r7gd.16xlarge": 4.354600, "r7gd.2xlarge": 0.544300, "r7gd.4xlarge": 1.088600,
		"r7gd.8xlarge": 2.177300, "r7gd.large": 0.136100, "r7gd.medium": 0.068000, "r7gd.metal": 4.354600,
		"r7gd.xlarge": 0.272200,
		// r7i family
		"r7i.12xlarge": 3.175200, "r7i.16xlarge": 4.233600, "r7i.24xlarge": 6.350400, "r7i.2xlarge": 0.529200,
		"r7i.48xlarge": 12.700800, "r7i.4xlarge": 1.058400, "r7i.8xlarge": 2.116800, "r7i.large": 0.132300,
		"r7i.metal-24xl": 6.985440, "r7i.metal-48xl": 12.700800, "r7i.xlarge": 0.264600,
		// r7iz family
		"r7iz.12xlarge": 4.464000, "r7iz.16xlarge": 5.952000, "r7iz.2xlarge": 0.744000, "r7iz.32xlarge": 11.904000,
		"r7iz.4xlarge": 1.488000, "r7iz.8xlarge": 2.976000, "r7iz.large": 0.186000, "r7iz.metal-16xl": 6.547200,
		"r7iz.metal-32xl": 13.094400, "r7iz.xlarge": 0.372000,
		// r8g family
		"r8g.12xlarge": 2.827680, "r8g.16xlarge": 3.770240, "r8g.24xlarge": 5.655360, "r8g.2xlarge": 0.471280,
		"r8g.48xlarge": 11.310720, "r8g.4xlarge": 0.942560, "r8g.8xlarge": 1.885120, "r8g.large": 0.117820,
		"r8g.medium": 0.058910, "r8g.metal-24xl": 6.220900, "r8g.metal-48xl": 11.310720, "r8g.xlarge": 0.235640,
		// t1 family
		"t1.micro": 0.020000,
		// t2 family
		"t2.2xlarge": 0.371200, "t2.large": 0.092800, "t2.medium": 0.046400, "t2.micro": 0.011600,
		"t2.nano": 0.005800, "t2.small": 0.023000, "t2.xlarge": 0.185600,
		// t3 family
		"t3.2xlarge": 0.332800, "t3.large": 0.083200, "t3.medium": 0.041600, "t3.micro": 0.010400,
		"t3.nano": 0.005200, "t3.small": 0.020800, "t3.xlarge": 0.166400,
		// t3a family
		"t3a.2xlarge": 0.300800, "t3a.large": 0.075200, "t3a.medium": 0.037600, "t3a.micro": 0.009400,
		"t3a.nano": 0.004700, "t3a.small": 0.018800, "t3a.xlarge": 0.150400,
		// t4g family
		"t4g.2xlarge": 0.268800, "t4g.large": 0.067200, "t4g.medium": 0.033600, "t4g.micro": 0.008400,
		"t4g.nano": 0.004200, "t4g.small": 0.016800, "t4g.xlarge": 0.134400,
		// trn1 family
		"trn1.2xlarge": 1.343750, "trn1.32xlarge": 21.500000,
		// trn1n family
		"trn1n.32xlarge": 24.780000,
		// u-12tb1 family
		"u-12tb1.112xlarge": 109.200000,
		// u-18tb1 family
		"u-18tb1.112xlarge": 163.800000,
		// u-24tb1 family
		"u-24tb1.112xlarge": 218.400000,
		// u-3tb1 family
		"u-3tb1.56xlarge": 27.300000,
		// u-6tb1 family
		"u-6tb1.112xlarge": 54.600000, "u-6tb1.56xlarge": 46.403910,
		// u-9tb1 family
		"u-9tb1.112xlarge": 81.900000,
		// u7i-12tb family
		"u7i-12tb.224xlarge": 152.880000,
		// u7i-6tb family
		"u7i-6tb.112xlarge": 62.790000,
		// u7i-8tb family
		"u7i-8tb.112xlarge": 83.720000,
		// u7in-16tb family
		"u7in-16tb.224xlarge": 203.840000,
		// u7in-24tb family
		"u7in-24tb.224xlarge": 305.760000,
		// u7in-32tb family
		"u7in-32tb.224xlarge": 407.680000,
		// vt1 family
		"vt1.24xlarge": 5.200000, "vt1.3xlarge": 0.650000, "vt1.6xlarge": 1.300000,
		// x1 family
		"x1.16xlarge": 6.669000, "x1.32xlarge": 13.338000,
		// x1e family
		"x1e.16xlarge": 13.344000, "x1e.2xlarge": 1.668000, "x1e.32xlarge": 26.688000, "x1e.4xlarge": 3.336000,
		"x1e.8xlarge": 6.672000, "x1e.xlarge": 0.834000,
		// x2gd family
		"x2gd.12xlarge": 4.008000, "x2gd.16xlarge": 5.344000, "x2gd.2xlarge": 0.668000, "x2gd.4xlarge": 1.336000,
		"x2gd.8xlarge": 2.672000, "x2gd.large": 0.167000, "x2gd.medium": 0.083500, "x2gd.metal": 5.878400,
		"x2gd.xlarge": 0.334000,
		// x2idn family
		"x2idn.16xlarge": 6.669000, "x2idn.24xlarge": 10.003500, "x2idn.32xlarge": 13.338000,
		"x2idn.metal": 13.338000,
		// x2iedn family
		"x2iedn.16xlarge": 13.338000, "x2iedn.24xlarge": 20.007000, "x2iedn.2xlarge": 1.667250,
		"x2iedn.32xlarge": 26.676000, "x2iedn.4xlarge": 3.334500, "x2iedn.8xlarge": 6.669000,
		"x2iedn.metal": 26.676000, "x2iedn.xlarge": 0.833630,
		// x2iezn family
		"x2iezn.12xlarge": 10.008000, "x2iezn.2xlarge": 1.668000, "x2iezn.4xlarge": 3.336000,
		"x2iezn.6xlarge": 5.004000, "x2iezn.8xlarge": 6.672000, "x2iezn.metal": 10.008000,
		// x8g family
		"x8g.12xlarge": 4.689600, "x8g.16xlarge": 6.252800, "x8g.24xlarge": 9.379200, "x8g.2xlarge": 0.781600,
		"x8g.48xlarge": 18.758400, "x8g.4xlarge": 1.563200, "x8g.8xlarge": 3.126400, "x8g.large": 0.195400,
		"x8g.medium": 0.097700, "x8g.metal-24xl": 10.317120, "x8g.metal-48xl": 18.758400, "x8g.xlarge": 0.390800,
		// z1d family
		"z1d.12xlarge": 4.464000, "z1d.2xlarge": 0.744000, "z1d.3xlarge": 1.116000, "z1d.6xlarge": 2.232000,
		"z1d.large": 0.186000, "z1d.metal": 4.464000, "z1d.xlarge": 0.372000,
	},
}
