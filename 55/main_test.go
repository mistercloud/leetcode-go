package main

import "testing"

func Test_canJump(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Test1",
			args{[]int{2, 3, 1, 1, 4}},
			true,
		},
		{
			"Test2",
			args{[]int{3, 2, 1, 0, 4}},
			false,
		},
		{
			"Test3",
			args{[]int{0, 2, 3}},
			false,
		},
		{
			"Test4",
			args{[]int{8049, 675, 4792, 2356, 16320, 15956, 9610, 7633, 13275, 10886, 11369, 18285, 1807, 12865, 6019, 13799, 16794, 7044, 10738, 11904, 7437, 14239, 18381, 8768, 12689, 2795, 566, 460, 16699, 7898, 6949, 17060, 15787, 1694, 1703, 19127, 13617, 19165, 13403, 6261, 17206, 14187, 16760, 3799, 9026, 9350, 19742, 19319, 16944, 8193, 9783, 12525, 13573, 5892, 16720, 16304, 13545, 12734, 6792, 7087, 9941, 1143, 9203, 10838, 7065, 10785, 7222, 355, 4915, 17279, 19005, 13193, 15373, 16199, 3505, 4902, 13793, 14143, 15942, 1863, 9896, 7651, 12331, 345, 16382, 8367, 6755, 6094, 8874, 9481, 1295, 7365, 3881, 6696, 10996, 18654, 506, 17960, 15474, 14971, 11084, 1437, 13610, 13017, 12470, 7079, 12880, 11243, 1852, 547, 17668, 14102, 5763, 7663, 8760, 3125, 1415, 16597, 3777, 2664, 13713, 16021, 15642, 3149, 9352, 5945, 5502, 6640, 1510, 6523, 7987, 8426, 10075, 15466, 444, 18990, 9878, 44, 4349, 7509, 12621, 16589, 18737, 18946, 6146, 12301, 4617, 11880, 947, 6240, 16845, 10496, 18061, 11403, 8224, 5828, 1814, 17207, 2279, 8655, 11254, 6763, 7248, 3406, 10827, 2212, 10350, 1196, 16382, 10298, 59, 5591, 7539, 18857, 835, 17487, 4704, 15038, 6837, 12452, 6121, 8666, 11243, 18401, 15227, 528, 15058, 12571, 4023, 19694, 19048, 1253, 10364, 9152, 3788, 9744, 13537, 4426, 11641, 17138, 1897, 18158, 18008, 5953, 18498, 16650, 9880, 17714, 4383, 4249, 6716, 18287, 16917, 17275, 552, 9987, 14953, 6722, 12328, 8987, 7528, 3663, 9484, 15283, 18398, 14436, 13516, 19754, 6116, 9232, 18452, 13587, 12096, 2964, 6597, 18551, 4489, 6428, 4997, 12045, 5690, 2989, 787, 7167, 11748, 4417, 7805, 6394, 18845, 47, 278, 714, 251, 19191, 3045, 18115, 12593, 3641, 14094, 14310, 4646, 12792, 5151, 10696, 9758, 17717, 13083, 17053, 3886, 4832, 18074, 18933, 3354, 15600, 2129, 11948, 15692, 9777, 12421, 17895, 3932, 7510, 2289, 2871, 17313, 5165, 16931, 5952, 12590, 16372, 9218, 17377, 16613, 9291, 1565, 19598, 14437, 7272, 6849, 986, 5660, 18559, 19897, 13657, 19117, 19154, 1967, 734, 7632, 5441, 3375, 17185, 19486, 9239, 287, 5224, 2171, 5751, 13338, 10429, 8107, 8832, 16728, 219, 8201, 2033, 2887, 19793, 17289, 9771, 8267, 19258, 11525, 8964, 19828, 13115, 2706, 17655, 15923, 9222, 14185, 1966, 5833, 9211, 14273, 19322, 4751, 6805, 10820, 107, 7169, 15538, 18235, 17094, 1437, 19011, 11928, 15896, 12197, 19944, 11123, 3801, 8986, 11016, 6048, 9199, 6065, 1137, 940, 12354, 9323, 2311, 18704, 15389, 10860, 18958, 6081, 3005, 15347, 17395, 19976, 13784, 6629, 14397, 12204, 19078, 11624, 15438, 18931, 11662, 6790, 12653, 17804, 12021, 844, 8125, 16643, 18943, 4326, 2490, 5032, 19924, 9135, 7103, 15166, 9680, 19519, 290, 2223, 864, 19389, 14194, 2109, 10658, 3495, 12845, 13035, 13244, 7369, 8384, 4308, 18154, 12053, 1413, 9647, 7863, 19471, 11498, 17075, 7094, 5626, 18893, 8255, 4773, 17303, 17258, 9641, 15252, 12418, 13663, 6173, 14853, 9377, 10921, 11787, 6797, 12773, 12034, 9269, 7237, 5770, 10281, 13242, 13601, 8034, 12674, 3704, 11896, 16471, 13351, 8524, 100, 4466, 13532, 9501, 17712, 9782, 13918, 12280, 1397, 1516, 19886, 4902, 10020, 662, 2170, 7786, 4713, 17788, 17524, 2632, 2726, 10965, 776, 4235, 10432, 10314, 8434, 887, 4228, 16375, 16739, 11451, 1013, 18933, 4015, 2391, 2954, 13049, 6918, 19876, 3069, 9194, 11225, 10511, 16214, 6790, 11269, 2725, 11723, 11712, 13126, 19034, 2497, 4136, 7684, 9938, 17071, 948, 9368, 10107, 10131, 12367, 12084, 18302, 3468, 7839, 10668, 19486, 10445, 3507, 375, 149, 14500, 2554, 10315, 15650, 2311, 9655, 3686, 2430, 16412, 13711, 3966, 6340, 16035, 11495, 2690, 14814, 7078, 6108, 17571, 12891, 18532, 7088, 4313, 3781, 2508, 9387, 11794, 11707, 14638, 8803, 18712, 2717, 15080, 8953, 4440, 18504, 7914, 7945, 10365, 9878, 5091, 18848, 12717, 10908, 7609, 17196, 10817, 15872, 19476, 5354, 15885, 5179, 14733, 14841, 120, 18246, 12493, 18368, 1548, 13548, 13561, 3023, 8882, 12142, 13659, 8672, 16288, 17301, 16449, 11396, 10525, 19278, 17930, 9118, 18758, 9607, 19907, 7018, 6245, 5675, 8503, 18816, 19184, 15328, 10762, 15307, 18706, 1925, 7591, 4400, 15063, 14510, 8750, 6220, 2778, 8802, 10079, 7036, 2202, 697, 14875, 1538, 13632, 17197, 2744, 2107, 4933, 12963, 16694, 6344, 17683, 962, 1240, 5360, 4782, 11184, 9415, 11672, 13455, 10485, 13780, 5795, 6642, 4339, 18015, 7214, 13432, 16859, 18863, 2496, 18778, 2496, 8678, 14281, 11222, 12489, 17926, 2832, 7942, 16398, 5084, 17818, 681, 4668, 13495, 7973, 978, 7096, 16785, 19548, 11774, 8825, 2475, 17464, 3596, 8207, 6197, 355, 11211, 10348, 17259, 18055, 8901, 13797, 12501, 6558, 2959, 8144, 18196, 991, 2495, 1123, 6193, 14656, 7262, 13080, 17327, 3470, 13572, 9836, 16651, 12795, 13496, 9246, 13629, 4186, 6045, 12247, 2122, 3791, 3329, 4290, 11641, 12128, 11569, 3084, 19747, 19253, 6804, 5277, 14255, 9033, 11496, 598, 13041, 13875, 16556, 13387, 3237, 17818, 189, 18336, 16701, 16253, 8322, 410, 19914, 15937, 2777, 3423, 17240, 14927, 3312, 6619, 10755, 7680, 7436, 13035, 2829, 14941, 15675, 6871, 14032, 15008, 12879, 3483, 15687, 16128, 5245, 18132, 3116, 2078, 11764, 3545, 3505, 11859, 9868, 4674, 14466, 6752, 9815, 3590, 18529, 11793, 13317, 8815, 7187, 17309, 13325, 12165, 2276, 7021, 7081, 4936, 10104, 11444, 9640, 732, 9775, 2690, 8255, 16443, 8733, 11575, 12612, 14059, 10229, 8866, 8696, 4130, 8216, 5728, 19762, 18917, 2995, 4261, 5092, 11061, 16660, 6425, 7192, 13140, 5759, 6433, 7037, 11879, 4168, 18092, 16289, 15415, 2742, 5460, 14228, 9424, 14033, 205, 13140, 4613, 7837, 19011, 16249, 16792, 4249, 17362, 14408, 747, 7369, 5639, 14417, 17504, 10156, 2758, 1459, 4725, 1101, 14625, 7078, 17028, 17873, 10335, 18115, 12203, 10836, 10902, 7939, 4357, 8217, 11153, 2538, 11172, 17293, 15402, 4068, 3382, 13909, 6754, 11161, 14268, 5448, 8347, 16596, 2216, 10157, 5211, 18542, 19490, 6175, 9499, 3720, 12158, 16529, 2155, 2910, 14784, 4946, 1774, 17233, 8537, 5272, 2973, 12012, 11596, 15030, 15008, 153, 2323, 8119, 18575, 13136, 9836, 9267, 17261, 18562, 14485, 10890, 13870, 4138, 14811, 1510, 4198, 15157, 19327, 13755, 10026, 4280, 11588, 15273, 2592, 5312, 7362, 141, 10725, 11321, 16846, 11421, 4006, 4186, 2285, 9296, 18093, 3219, 4184, 18207, 1853, 17294, 19082, 8018, 13172, 6544, 4736, 14453, 15398, 8887, 2285, 4569, 1842, 12660, 10900, 18251, 12974, 10706, 16392, 2182, 11513, 13402, 3488, 3205, 495, 1394, 8670, 13911, 11735, 14795, 17102, 1114, 18405, 767, 17791, 483, 5997, 13720, 5571, 3657, 8720, 9146, 17745, 315, 5745, 5379, 14377, 16607, 17760, 11978, 12799, 12779, 13384, 3628, 6000, 17726, 1806, 5961, 17519, 6419, 12239, 9383, 12089, 7699, 1763, 12503, 17433, 4882, 13971, 1252, 7465, 7123, 17976, 16136, 5658, 10028, 18005, 7340, 4269, 2784, 13961, 19024, 2695, 493, 11894, 6047, 12937, 15394, 15971, 5901, 5321, 19829, 9974, 8549, 16753, 11901, 7188, 18370, 5181, 15540, 9361, 18667, 10542, 4022, 788, 3103, 7835, 6138, 14780, 14860, 5707, 18993, 16220, 15246, 5953, 3476, 17902, 17073, 15448, 17405, 10843, 16159, 7617, 934, 9885, 11034, 9320, 9495, 11407, 415, 10678, 12793, 4479, 10011, 8952, 841, 6262, 6836, 1152, 2054, 17521, 10354, 17133, 12649, 6412, 9540, 2893, 58, 9766, 5530, 18172, 2225, 4129, 4665, 6829, 1186, 13698, 11480, 19559, 1497, 13408, 12319, 10308, 17543, 6643, 4177, 4029, 14404, 14439, 1142, 3558, 10473, 16091, 13156, 582, 10185, 5022, 13729, 6347, 222, 4442, 12997, 18541, 6421, 5100, 19164, 10795, 1868, 12581, 9244, 11009, 5033, 2028, 9187, 6360, 6574, 18858, 16099, 14950, 2798, 10860, 18959, 12358, 14308, 4496, 3342, 755, 18929, 18775, 13962, 8960, 19269, 3542, 3452, 6390, 4884, 17616, 12539, 15272, 17299, 11496, 3458, 3092, 2323, 12296, 17529, 9150, 14854, 10871, 13251, 12925, 18991, 300, 14695, 5551, 11964, 15116, 8803, 11445, 8060, 1061, 8951, 6779, 725, 14429, 7553, 16838, 14915, 10518, 11202, 712, 13928, 116, 10720, 2838, 12734, 19204, 6348, 357, 731, 6156, 16017, 17564, 15063, 6526, 8304, 6654, 15931, 2781, 9586, 5482, 12410, 3349, 812, 9459, 19013, 8315, 498, 4245, 3572, 19870, 7836, 10846, 6620, 9479, 2920, 1503, 8114, 12069, 1359, 7512, 10091, 9966, 3749, 9374, 19034, 5978, 6447, 11286, 16740, 4001, 6347, 17763, 15642, 15277, 17298, 5908, 4207, 6141, 17333, 3407, 13377, 678, 5284, 11542, 8124, 14100, 3299, 883, 13392, 12077, 3619, 12363, 8683, 8660, 4627, 5638, 14962, 17849, 12583, 1812, 7338, 12344, 11857, 3518, 5678, 3473, 11884, 7682, 12467, 521, 15410, 8783, 11467, 2501, 5774, 5497, 14979, 8100, 11761, 12264, 17716, 2026, 804, 7487, 11400, 14710, 16877, 14480, 6773, 8470, 19575, 7865, 1169, 16898, 12814, 14825, 15983, 18767, 4623, 18616, 11175, 10353, 5775, 1846, 13269, 16860, 6232, 6985, 12605, 11658, 19068, 7539, 11842, 207, 364, 11841, 13710, 14509, 9639, 15472, 17984, 10303, 8964, 16954, 1054, 15412, 15500, 13532, 18336, 18671, 6879, 18449, 15536, 14397, 7793, 17419, 8964, 14315, 13252, 718, 13949, 171, 10816, 10141, 3113, 12341, 17257, 14855, 14456, 18008, 2010, 3232, 10194, 7238, 15438, 10250, 8139, 9639, 18274, 9012, 7822, 12794, 17577, 15601, 19785, 4428, 15214, 6996, 15058, 8742, 4918, 10784, 8272, 11307, 9307, 8772, 4069, 5800, 18852, 3579, 16819, 19167, 6160, 6860, 883, 11330, 19922, 5024, 11148, 1447, 7108, 972, 11209, 2871, 466, 16235, 15038, 18704, 3839, 286, 13732, 1427, 15840, 1138, 460, 10485, 18704, 11383, 3786, 9814, 15043, 15806, 1776, 5799, 18506, 4819, 13447, 6871, 1090, 1663, 6005, 10292, 14538, 19144, 482, 16053, 6797, 3338, 18325, 2560, 8796, 5243, 404, 10017, 13738, 11142, 3551, 51, 16707, 7202, 13892, 2540, 13721, 7340, 18743, 10866, 18035, 2201, 12139, 17968, 11960, 4156, 18787, 6961, 9099, 13718, 9144, 2174, 17410, 17592, 17470, 514, 8860, 5978, 5003, 16454, 11309, 16322, 19252, 9953, 7797, 5206, 1967, 4518, 9620, 19484, 5804, 15969, 8398, 11074, 17215, 3379, 1277, 10025, 3438, 13996, 17040, 14874, 14170, 8663, 5735, 11698, 7615, 6169, 15823, 914, 1873, 16082, 6878, 4393, 14814, 16302, 488, 6267, 19824, 14301, 2534, 17149, 7560, 10560, 17425, 10445, 5370, 4205, 6539, 12500, 10264, 9283, 6126, 9098, 4095, 12984, 17473, 10033, 19184, 7237, 8939, 17208, 4687, 1532, 10381, 7166, 8226, 5483, 1924, 168, 4583, 19842, 8247, 10788, 3587, 17222, 10132, 11239, 12848, 1339, 19881, 14592, 7499, 14569, 4116, 11790, 13911, 11134, 9822, 712, 8712, 6932, 4160, 16242, 8197, 14987, 8971, 7409, 10073, 9958, 6690, 13475, 3718, 6556, 8273, 2170, 15271, 12279, 8825, 10213, 11983, 19539, 13173, 7114, 2632, 8749, 4940, 17256, 7091, 3294, 4685, 1564, 9858, 16421, 1710, 13913, 6319, 13267, 13073, 472, 7275, 13521, 8103, 15516, 2145, 16374, 10806, 18626, 8013, 1402, 2283, 7535, 17373, 15531, 18926, 4659, 221, 18841, 4845, 2962, 12948, 5977, 9032, 17350, 10002, 2866, 17291, 4761, 13756, 5454, 5165, 18210, 8244, 3294, 3883, 14600, 12591, 18728, 6193, 3968, 18634, 11640, 11752, 2302, 7364, 10272, 3773, 758, 15084, 8394, 19643, 18600, 17203, 6748, 16678, 16571, 4189, 3902, 12293, 597, 1276, 4625, 7029, 4685, 1805, 2930, 3518, 6441, 15723, 9906, 120, 14584, 6087, 15857, 6403, 19881, 10214, 12171, 8023, 11648, 12343, 8818, 14014, 12560, 1134, 11748, 15403, 17005, 1378, 10198, 8990, 12715, 12147, 13640, 19642, 7607, 19648, 11970, 11002, 11282, 16752, 10578, 12264, 2803, 14636, 13378, 16116, 7392, 5940, 3801, 7819, 13483, 14166, 7795, 10847, 12812, 14665, 17973, 10535, 17877, 17733, 16755, 1604, 4583, 6247, 15062, 7432, 12072, 15555, 11708, 7504, 4705, 17229, 8770, 2101, 16978, 1798, 8553, 12810, 17729, 18373, 9676, 18342, 166, 16343, 9275, 15801, 17970, 10237, 11541, 13420, 16673, 5633, 7859, 19696, 13289, 12577, 12140, 8337, 9218, 15719, 10822, 10919, 4143, 19410, 9734, 17503, 10369, 18189, 3420, 8225, 10237, 2524, 10237, 11344, 8776, 17300, 18582, 17255, 10326, 345, 7345, 18258, 15895, 16819, 19758, 11803, 10649, 12529, 3620, 292, 3597, 3794, 18085, 6329, 19675, 19234, 5708, 15949, 7876, 1529, 1711, 5329, 1171, 372, 2467, 2103, 18650, 13604, 41, 3499, 12268, 13563, 14155, 8155, 5616, 15818, 4348, 8338, 9589, 19383, 10804, 16766, 11550, 7119, 11566, 19782, 14074, 10813, 1403, 17693, 1051, 2269, 14819, 2479, 10990, 5468, 1256, 2452, 8002, 10023, 3042, 16186, 16379, 11121, 17773, 16294, 11546, 10680, 326, 18663, 12263, 10653, 5507, 2543, 6021, 9918, 12568, 9138, 18838, 8851, 19556, 17893, 1077, 1245, 18444, 15564, 16568, 770, 17089, 1493, 10454, 15081, 10112, 2906, 19776, 8727, 11237, 10963, 4716, 12104, 17937, 14818, 1043, 4756, 17859, 15330, 9323, 3274, 4433, 5489, 17796, 17152, 7140, 14563, 17798, 17818, 3960, 19073, 4913, 6605, 9464, 7741, 19713, 9220, 7966, 13170, 1555, 6230, 8961, 2137, 6692, 209, 12892, 17986, 14457, 14994, 10501, 16925, 16221, 858, 8311, 2300, 15870, 12141, 13994, 17227, 6481, 8765, 4775, 10447, 15177, 13009, 6883, 11250, 6592, 6524, 17357, 4722, 19037, 1955, 12103, 2565, 7753, 3801, 12226, 10749, 5200, 18688, 18293, 14638, 8059, 2881, 14746, 13522, 19915, 7293, 6188, 17887, 4769, 3589, 18370, 9245, 7392, 5258, 2227, 1818, 3759, 7836, 973, 9210, 10799, 16743, 923, 346, 16849, 9553, 12072, 19392, 10598, 5523, 14449, 3880, 11835, 5613, 6295, 6493, 8139, 9302, 14872, 12421, 1978, 10252, 16872, 16998, 3736, 6896, 15068, 19061, 8610, 10258, 12631, 4269, 7084, 19537, 2174, 17451, 18753, 1821, 15518, 12704, 3517, 3040, 9768, 8222, 2732, 4780, 13843, 8037, 399, 15150, 174, 8014, 12814, 2483, 19105, 14256, 4151, 9737, 12593, 2837, 16816, 179, 15038, 422, 5945, 8975, 14781, 2114, 5100, 12461, 1456, 3621, 14257, 12413, 19508, 16766, 16693, 11656, 3180, 1106, 8442, 17909, 17128, 999, 10900, 1640, 13339, 8104, 10911, 2926, 7054, 15001, 3854, 16871, 3786, 16679, 16175, 8120, 307, 13835, 8602, 7801, 6475, 8043, 9063, 3052, 9437, 1020, 18315, 10946, 12133, 11986, 3777, 10725, 65, 19608, 267, 12279, 5793, 18061, 12302, 6229, 5008, 11489, 13626, 8270, 7765, 3926, 4875, 19140, 14003, 5058, 10115, 19597, 2250, 2455, 13232, 4803, 18283, 7856, 16054, 12473, 13751, 17118, 18463, 618, 12787, 13669, 12567, 4872, 12489, 4875, 17148, 12849, 8269, 19854, 18539, 6027, 3177, 18466, 819, 11298, 16394, 18525, 14852, 16250, 16952, 8450, 8891, 14739, 3483, 9857, 4041, 9448, 15633, 14297, 13642, 15167, 10176, 11840, 15175, 1583, 6455, 10179, 6467, 11767, 7175, 4206, 2086, 16835, 13370, 12479, 11640, 12267, 8900, 7517, 17511, 3192, 6489, 11557, 14353, 2323, 16120, 14140, 15966, 3869, 16420, 17523, 5399, 17276, 4078, 19994, 9119, 5126, 1712, 6286, 13193, 14524, 18706, 11936, 18028, 1036, 473, 2929, 5592, 15263, 8140, 14973, 16410, 9785, 8705, 11098, 8487, 5760, 19829, 16557, 16422, 2560, 11669, 18884, 5177, 3558, 18984, 17357, 11626, 4862, 5181, 9883, 8022, 6524, 16713, 13197, 12075, 289, 5335, 2332, 7149, 5666, 8480, 11333, 6074, 8279, 601, 9955, 1658, 19996, 15663, 19881, 10467, 15886, 14338, 8536, 11349, 16834, 10849, 12952, 6987, 17598, 12696, 4461, 3259, 10221, 7838, 2983, 8645, 7414, 9150, 1252, 15430, 10059, 19813, 194, 2995, 13377, 8095, 3795, 3063, 11097, 19197, 4291, 6114, 11356, 7587, 4199, 8938, 2178, 4418, 8010, 502, 1227, 14838, 14099, 7638, 5895, 18902, 3097, 8085, 2949, 8439, 9734, 16303, 19915, 12232, 9375, 13681, 8586, 1007, 8820, 2370, 10117, 9688, 14256, 9786, 6005, 14028, 992, 4450, 13272, 14223, 458, 19651, 8904, 11912, 11461, 15139, 14888, 17172, 14977, 15655, 4177, 4159, 17762, 10721, 1108, 10063, 18434, 19111, 6476, 12919, 7684, 17948, 12043, 3524, 14915, 16792, 11446, 17695, 3233, 638, 4288, 13198, 3113, 9420, 8614, 19912, 9210, 692, 4566, 7758, 7601, 16224, 15590, 4314, 5631, 1956, 7320, 1694, 15029, 13224, 15167, 5037, 6903, 13029, 13851, 12961, 3275, 11063, 937, 17364, 16592, 19005, 5325, 9684, 15819, 4904, 9432, 15431, 15268, 10615, 11617, 12649, 3477, 18648, 19183, 6400, 7714, 8628, 12646, 5890, 17317, 5477, 9750, 1607, 9504, 9319, 19008, 2160, 12419, 13169, 4626, 16146, 18597, 6021, 8985, 18202, 9927, 16269, 14540, 2737, 6700, 7171, 14230, 2823, 3063, 18294, 14906, 3847, 1581, 714, 1501, 4089, 10073, 10608, 1525, 6679, 9064, 4208, 2839, 13850, 1047, 9845, 7290, 10198, 16219, 2181, 2459, 7513, 5961, 5371, 5297, 16999, 434, 19143, 1006, 2770, 9514, 12246, 12657, 17351, 9934, 18426, 8429, 14273, 14293, 16449, 19586, 268, 8690, 529, 5759, 13312, 15049, 8115, 6142, 2666, 13713, 17399, 7443, 2003, 15278, 1676, 16978, 8815, 15231, 17734, 19547, 6582, 3730, 11398, 15733, 12849, 17423, 15147, 2496, 3952, 5799, 17891, 12505, 12520, 2243, 18944, 8737, 3132, 2992, 48, 17414, 1908, 19619, 14624, 12056, 15846, 15234, 12177, 12366, 420, 17256, 9135, 5716, 11037, 12722, 19456, 16517, 10628, 13573, 10584, 12341, 11855, 8842, 12819, 16412, 5520, 15725, 18354, 3153, 18207, 733, 5742, 15647, 13279, 8807, 6955, 934, 9342, 11256, 14692, 19570, 5911, 19336, 16896, 2407, 3256, 15817, 6852, 18998, 11940, 5899, 4481, 16135, 6234, 13351, 47, 4419, 11090, 3182, 13745, 16148, 892, 16202, 10720, 9176, 8988, 17856, 585, 17400, 694, 4136, 9396, 6301, 4031, 9299, 1657, 14820, 19102, 19779, 1107, 16297, 9187, 10253, 2177, 5687, 2265, 8496}},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canJump(tt.args.nums); got != tt.want {
				t.Errorf("canJump() = %v, want %v", got, tt.want)
			}
		})
	}
}
