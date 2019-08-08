package io.viascom.vleas.models.mavencentral

import java.io.Serializable

class Params(
    val q: String,
    val core: String,
    val indent: String,
    val fl: String,
    val start: String,
    val sort: String,
    val rows: String,
    val wt: String,
    val version: String
) : Serializable 
