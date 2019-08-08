package io.viascom.vleas.models.mavencentral

import java.io.Serializable

class Doc(
    val id: String,
    val g: String,
    val a: String,
    val v: String,
    val p: String,
    val timestamp: String,
    val ec: List<String>,
    val tags: List<String>
) : Serializable