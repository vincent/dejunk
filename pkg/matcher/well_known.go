package matcher

import "regexp"

// ItemType holds
type ItemType string

const (
	// Music ItemType
	Music ItemType = "Music"
	// Movie ItemType
	Movie = "Movie"
	// TVShow ItemType
	TVShow = "TVShow"
	// Other ItemType
	Other = "Other"
)

// SeasonEpisode is RESeasonEpisode regex string
var SeasonEpisode = `([Ss]?)(\d{1,2})([xXeE\.\-])(\d{1,6})`

// RESeasonEpisode match SxEy
var RESeasonEpisode = regexp.MustCompile(`([Ss]?)(\d{1,2})([xXeE\.\-])(\d{1,6})`)

// REYear match YYYY
var REYear = regexp.MustCompile(`[^\d](\d{4})([^\d]|$)`)

// RECleanJunk match non alpha chars
var RECleanJunk = regexp.MustCompile(`[-'~!@#$%^&*()_|+=?;:'",.<>\{\}\[\]\\\/]`)

// RECleanSpaces match multispaces
var RECleanSpaces = regexp.MustCompile(`\s+`)

// AudioFileExts holds all known audio file formats
var AudioFileExts = []string{
	".3gp",   //  Multimedia container format can contain proprietary formats as AMR, AMR-WB or AMR-WB+, but also some open formats
	".aa",    //  Audible (Amazon.com)  A low-bitrate audiobook container format with DRM, containing audio encoded as either MP3 or the ACELP speech codec.
	".aac",   //  The Advanced Audio Coding format is based on the MPEG-2 and MPEG-4 standards. AAC files are usually ADTS or ADIF containers.
	".aax",   //  Audible (Amazon.com)  An Audiobook format, which is a variable-bitrate (allowing high quality) M4B file encrypted with DRM. MPB contains AAC or ALAC encoded audio in an MPEG-4 container. (More details below.)
	".act",   //  ACT is a lossy ADPCM 8 kbit/s compressed audio format recorded by most Chinese MP3 and MP4 players with a recording function, and voice recorders
	".aiff",  //  Apple  A standard uncompressed CD-quality, audio file format used by Apple. Established 3 years prior to Microsoft's uncompressed version wav.
	".alac",  //  Apple  An audio coding format developed by Apple Inc. for lossless data compression of digital music.
	".amr",   //  AMR-NB audio, used primarily for speech.
	".ape",   //  Matthew T. Ashland  Monkey's Audio lossless audio compression format.
	".au",    //  Sun Microsystems  The standard audio file format used by Sun, Unix and Java. The audio in au files can be PCM or compressed with the μ-law, a-law or G729 codecs.
	".awb",   //  AMR-WB audio, used primarily for speech, same as the ITU-T's G.722.2 specification.
	".dct",   //  NCH Software  A variable codec format designed for dictation. It has dictation header information and can be encrypted (as may be required by medical confidentiality laws). A proprietary format of NCH Software.
	".dss",   //  Olympus  DSS files are an Olympus proprietary format. It is a fairly old and poor codec. GSM or MP3 are generally preferred where the recorder allows. It allows additional data to be held in the file header.
	".dvf",   //  Sony  A Sony proprietary format for compressed voice files; commonly used by Sony dictation recorders.
	".flac",  //  A file format for the Free Lossless Audio Codec, an open-source lossless compression codec.
	".gsm",   //  Designed for telephony use in Europe, gsm is a very practical format for telephone quality voice. It makes a good compromise between file size and quality. Note that wav files can also be encoded with the gsm codec.
	".iklax", //  iKlax  An iKlax Media proprietary format, the iKlax format is a multi-track digital audio format allowing various actions on musical data, for instance on mixing and volumes arrangements.
	".ivs",   //  3D Solar UK Ltd  A proprietary version with Digital Rights Management developed by 3D Solar UK Ltd for use in music downloaded from their Tronme Music Store and interactive music and video player.
	".m4a",   //  An audio-only MPEG-4 file, used by Apple for unprotected music downloaded from their iTunes Music Store. Audio within the m4a file is typically encoded with AAC, although lossless ALAC may also be used.
	".m4b",   //  An Audiobook / podcast extension with AAC or ALAC encoded audio in an MPEG-4 container. Both M4A and M4B formats can contain metadata including chapter markers, images, and hyperlinks, but M4B allows "bookmarks" (remembering the last listening spot), whereas M4A does not.[7]
	".m4p",   //  Apple  A version of AAC with proprietary Digital Rights Management developed by Apple for use in music downloaded from their iTunes Music Store.
	".mmf",   //  Yamaha, Samsung  A Samsung audio format that is used in ringtones. Developed by Yamaha (SMAF stands for "Synthetic music Mobile Application Format", and is a multimedia data format invented by the Yamaha Corporation, .mmf file format).
	".mp3",   //  MPEG Layer III Audio. It is the most common sound file format used today.
	".mpc",   //  Musepack or MPC (formerly known as MPEGplus, MPEG+ or MP+) is an open source lossy audio codec, specifically optimized for transparent compression of stereo audio at bitrates of 160–180 kbit/s.
	".msv",   //  Sony  A Sony proprietary format for Memory Stick compressed voice files.
	".nmf",   //  NICE  NICE Media Player audio file
	".ogg",   //, .oga, .mogg  Xiph.Org Foundation  A free, open source container format supporting a variety of formats, the most popular of which is the audio format Vorbis. Vorbis offers compression similar to MP3 but is less popular. Mogg, the "Multi-Track-Single-Logical-Stream Ogg-Vorbis", is the multi-channel or multi-track Ogg file format.
	".opus",  //  Internet Engineering Task Force  A lossy audio compression format developed by the Internet Engineering Task Force (IETF) and made especially suitable for interactive real-time applications over the Internet. As an open format standardised through RFC 6716, a reference implementation is provided under the 3-clause BSD license.
	".org",   //  Daisuke
	".ra",    //, .rm  RealNetworks  A RealAudio format designed for streaming audio over the Internet. The .ra format allows files to be stored in a self-contained fashion on a computer, with all of the audio data contained inside the file itself.
	".raw",   //  A raw file can contain audio in any format but is usually used with PCM audio data. It is rarely used except for technical tests.
	".rf64",  //  One successor to the Wav format, overcoming the 4GiB size limitation.
	".sln",   //  Signed Linear PCM format used by Asterisk. Prior to v.10 the standard formats were 16-bit Signed Linear PCM sampled at 8 kHz and at 16 kHz. With v.10 many more sampling rates were added.[8]
	".tta",   //  The True Audio, real-time lossless audio codec.
	".voc",   //  Creative Technology  The file format consists of a 26-byte header and a series of subsequent data blocks containing the audio information
	".vox",   //  The vox format most commonly uses the Dialogic ADPCM (Adaptive Differential Pulse Code Modulation) codec. Similar to other ADPCM formats, it compresses to 4-bits. Vox format files are similar to wave files except that the vox files contain no information about the file itself so the codec sample rate and number of channels must first be specified in order to play a vox file.
	".wav",   //  Standard audio file container format used mainly in Windows PCs. Commonly used for storing uncompressed (PCM), CD-quality sound files, which means that they can be large in size—around 10 MB per minute. Wave files can also contain data encoded with a variety of (lossy) codecs to reduce the file size (for example the GSM or MP3 formats). Wav files use a RIFF structure.
	".wma",   //  Microsoft  Windows Media Audio format, created by Microsoft. Designed with Digital Rights Management (DRM) abilities for copy protection.
	".wv",    //  Format for wavpack files.
	".webm",  //  Royalty-free format created for HTML5 video.
	".8svx",  //  Electronic Arts  The IFF-8SVX format for 8-bit sound samples, created by Electronic Arts in 1984 at the birth of the Amiga.
	".cda",   //  Format for cda files for Radio.
}

// VideoFileExts holds all known video file formats
var VideoFileExts = []string{
	".webm", // WebM Matroska  VP8, VP9, AV1  Vorbis, Opus  Royalty-free format created for HTML5 video.
	".mkv",  // Matroska Matroska  any  any
	".flv",  // Flash Video (FLV) FLV  VP6, Sorenson Spark, Screen video, Screen video 2, H.264  MP3, ADPCM, Nellymoser, Speex, AAC  Use of the H.264 and AAC compression formats in the FLV file format has some limitations and authors of Flash Player strongly encourage everyone to embrace the new standard F4V file format[2] de facto standard for web-based streaming video (over RTMP).
	".flv",  // F4V MPEG-4 Part 12  H.264  MP3, AAC  Replacement for FLV.
	".vob",  // Vob VOB  H.262/MPEG-2 Part 2 or MPEG-1 Part 2  PCM, DTS, MPEG-1, Audio Layer II (MP2), or Dolby Digital (AC-3)  Files in VOB format have .vob filename extension and are typically stored in the VIDEO_TS folder at the root of a DVD. The VOB format is based on the MPEG program stream format.
	".ogv",  // Ogg Video, .ogg  Ogg  Theora, Dirac  Vorbis, FLAC
	".drc",  // Dirac Dirac  ?
	".gif",  // GIF N/A  N/A  none  Simple animation, inefficient compression, no sound, widely supported
	".gifv", // Video alternative to GIF HTML  any  none  Not standardized, and not a real video file in the classical meaning since it merely references the real video file (e.g. a .webm file), which has to exist separately elsewhere. A .gifv "file" is simply a HTML webpage which includes a HTML5 video tag, where the video has no sound. As there were large communities online which create art using the medium of short soundless videos in GIF format, GIFV was created as a functionally similar replacement with vastly smaller filesizes than the inefficient GIF format.
	".mng",  // Multiple-image Network Graphics N/A  N/A  none  Inefficient, not widely used.
	".avi",  // AVI AVI  any  any  Uses RIFF
	".MTS",  // MPEG Transport Stream, .M2TS, .TS  AVCHD  AVCHD (MPEG-4 / H.264 )  Dolby AC-3 or uncompressed linear PCM  The standard video format used by many Sony and Panasonic HD camcorders. It is also used for storing high definition video on Blu-ray discs.
	".mov",  // QuickTime File Format, .qt  QuickTime  many[3]  AAC, MP3, others[3]
	".wmv",  // Windows Media Video ASF  Windows Media Video, Windows Media Video Screen, Windows Media Video Image  Windows Media Audio, Sipro ACELP.net
	".yuv",  // Raw video format Further documentation needed  Doesn't apply  Doesn't apply  Supports all resolutions, sampling structures, and frame rates
	".rm",   // RealMedia (RM) RealMedia  RealVideo  RealAudio  Made for RealPlayer
	".rmvb", // RealMedia Variable Bitrate (RMVB) RealMedia Variable Bitrate  RealVideo  RealAudio  Made for RealPlayer
	".viv",  // VivoActive (VIV) VIV  based upon H.263 video  G.723 ADPCM audio (not the G.723.1 speech codec)  Made for VivoActive Player
	".asf",  // Advanced Systems Format (ASF) ASF  any  any
	".amv",  // AMV video format Modified version of AVI[4]  Variant of Motion JPEG  Variant of IMA, ADPCM  Proprietary video file format produced for MP4 players and S1 MP3 players with video playback
	".mp4",  // MPEG-4 Part 14 (MP4), .m4p (with DRM), .m4v  MPEG-4 Part 12  H.264, MPEG-4 Part 2, MPEG-2, MPEG-1  Advanced Audio Coding, MP3, others
	".mpg",  // MPEG-1, .mp2, .mpeg, .mpe, .mpv  MPEG-1 part 1  MPEG-1 part 2  MPEG-1 Audio Layer I, MPEG-1 Audio Layer I, MPEG-1 Audio Layer III (MP3)  Old, but very widely used due to installed base.
	".mpg",  // MPEG-2 – Video, .mpeg, .m2v  H.262  AAC, MP3, MPEG-2 Part 3, others
	".m4v",  // M4V – (file format for videos for iPods and PlayStation Portables developed by Apple) MPEG-4 Part 12  H.264  AAC, Dolby Digital  Developed by Apple, used in iTunes. Very similar to MP4 format, but may optionally have DRM.
	".svi",  // SVI MPEG-4 utilising a special header   Samsung video format for portable players
	".3gp",  // 3GPP MPEG-4 Part 12  MPEG-4 Part 2, H.263, H.264  AMR-NB, AMR-WB, AMR-WB+, AAC-LC, HE-AAC v1 or Enhanced aacPlus (HE-AAC v2)  Common video format for cell phones
	".3g2",  // 3GPP2 MPEG-4 Part 12  MPEG-4 Part 2, H.263, H.264  AMR-NB, AMR-WB, AMR-WB+, AAC-LC, HE-AAC v1 or Enhanced aacPlus (HE-AAC v2), EVRC, SMV or VMR-WB  Common video format for cell phones
	".mxf",  // Material Exchange Format (MXF) MXF
	".roq",  // ROQ  used by Quake 3[5]
	".nsv",  // Nullsoft Streaming Video (NSV) NSV   For streaming video content over the Internet
	".flv",  // Flash Video (FLV)   .f4v .f4p .f4a .f4b  Audio, video, text, data  Adobe Flash Platform  SWF, F4V, ISO base media file format  Developed by the Adobe Flash Platform
}
