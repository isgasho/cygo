package dwarf

func TagKindName(kind int) string {
	name := TagNames[kind]
	if name == "" {
		name = "tagunk-" + kind.repr()
	} else {
		name = name.trimsp()
	}
	return name
}

func AttrKindName(kind int) string {
	// name, ok := AttrNames[kind] // TODO compiler
	name := AttrNames[kind]
	if name == "" {
		name = "attrunk-" + kind.repr()
	} else {
		name = name.trimsp()
	}
	return name
}

func FormKindName(kind int) string {
	name := FormNames[kind]
	if name == "" {
		name = "formunk-" + kind.repr()
	} else {
		name = name.trimsp()
	}
	return name
}

func ErrorKindName(kind int) string {
	name := ErrorNames[kind]
	if name == "" {
		name = "errunk-" + kind.repr()
	} else {
		name = name.trimsp()
	}
	return name
}

var (
	TagNames  = map[int]string{}
	AttrNames = map[int]string{
		DW_AT_sibling:          "DW_AT_sibling                 ",
		DW_AT_location:         "DW_AT_location                ",
		DW_AT_name:             "DW_AT_name                    ",
		DW_AT_ordering:         "DW_AT_ordering                ",
		DW_AT_subscr_data:      "DW_AT_subscr_data             ",
		DW_AT_byte_size:        "DW_AT_byte_size               ",
		DW_AT_bit_offset:       "DW_AT_bit_offset              ",
		DW_AT_bit_size:         "DW_AT_bit_size                ",
		DW_AT_element_list:     "DW_AT_element_list            ",
		DW_AT_stmt_list:        "DW_AT_stmt_list               ",
		DW_AT_low_pc:           "DW_AT_low_pc                  ",
		DW_AT_high_pc:          "DW_AT_high_pc                 ",
		DW_AT_language:         "DW_AT_language                ",
		DW_AT_member:           "DW_AT_member                  ",
		DW_AT_discr:            "DW_AT_discr                   ",
		DW_AT_discr_value:      "DW_AT_discr_value             ",
		DW_AT_visibility:       "DW_AT_visibility              ",
		DW_AT_import:           "DW_AT_import                  ",
		DW_AT_string_length:    "DW_AT_string_length           ",
		DW_AT_common_reference: "DW_AT_common_reference        ",
		DW_AT_comp_dir:         "DW_AT_comp_dir                ",
		DW_AT_const_value:      "DW_AT_const_value             ",
		DW_AT_containing_type:  "DW_AT_containing_type         ",
		DW_AT_default_value:    "DW_AT_default_value           ",
		DW_AT_inline:           "DW_AT_inline                  ",
		DW_AT_is_optional:      "DW_AT_is_optional             ",
		DW_AT_lower_bound:      "DW_AT_lower_bound             ",
		DW_AT_producer:         "DW_AT_producer                ",
		DW_AT_prototyped:       "DW_AT_prototyped              ",
		DW_AT_return_addr:      "DW_AT_return_addr             ",
		DW_AT_start_scope:      "DW_AT_start_scope             ",
		DW_AT_bit_stride:       "DW_AT_bit_stride              ",
		// DW_AT_stride_size:             "DW_AT_stride_size             ",
		DW_AT_upper_bound:          "DW_AT_upper_bound             ",
		DW_AT_abstract_origin:      "DW_AT_abstract_origin         ",
		DW_AT_accessibility:        "DW_AT_accessibility           ",
		DW_AT_address_class:        "DW_AT_address_class           ",
		DW_AT_artificial:           "DW_AT_artificial              ",
		DW_AT_base_types:           "DW_AT_base_types              ",
		DW_AT_calling_convention:   "DW_AT_calling_convention      ",
		DW_AT_count:                "DW_AT_count                   ",
		DW_AT_data_member_location: "DW_AT_data_member_location    ",
		DW_AT_decl_column:          "DW_AT_decl_column             ",
		DW_AT_decl_file:            "DW_AT_decl_file               ",
		DW_AT_decl_line:            "DW_AT_decl_line               ",
		DW_AT_declaration:          "DW_AT_declaration             ",
		DW_AT_discr_list:           "DW_AT_discr_list              ",
		DW_AT_encoding:             "DW_AT_encoding                ",
		DW_AT_external:             "DW_AT_external                ",
		DW_AT_frame_base:           "DW_AT_frame_base              ",
		DW_AT_friend:               "DW_AT_friend                  ",
		DW_AT_identifier_case:      "DW_AT_identifier_case         ",
		DW_AT_macro_info:           "DW_AT_macro_info              ",
		DW_AT_namelist_item:        "DW_AT_namelist_item           ",
		DW_AT_priority:             "DW_AT_priority                ",
		DW_AT_segment:              "DW_AT_segment                 ",
		DW_AT_specification:        "DW_AT_specification           ",
		DW_AT_static_link:          "DW_AT_static_link             ",
		DW_AT_type:                 "DW_AT_type                    ",
		DW_AT_use_location:         "DW_AT_use_location            ",
		DW_AT_variable_parameter:   "DW_AT_variable_parameter      ",
		DW_AT_virtuality:           "DW_AT_virtuality              ",
		DW_AT_vtable_elem_location: "DW_AT_vtable_elem_location    ",
		DW_AT_allocated:            "DW_AT_allocated               ",
		DW_AT_associated:           "DW_AT_associated              ",
		DW_AT_data_location:        "DW_AT_data_location           ",
		DW_AT_byte_stride:          "DW_AT_byte_stride             ",
		// DW_AT_stride:                  "DW_AT_stride                  ",
		DW_AT_entry_pc:                "DW_AT_entry_pc                ",
		DW_AT_use_UTF8:                "DW_AT_use_UTF8                ",
		DW_AT_extension:               "DW_AT_extension               ",
		DW_AT_ranges:                  "DW_AT_ranges                  ",
		DW_AT_trampoline:              "DW_AT_trampoline              ",
		DW_AT_call_column:             "DW_AT_call_column             ",
		DW_AT_call_file:               "DW_AT_call_file               ",
		DW_AT_call_line:               "DW_AT_call_line               ",
		DW_AT_description:             "DW_AT_description             ",
		DW_AT_binary_scale:            "DW_AT_binary_scale            ",
		DW_AT_decimal_scale:           "DW_AT_decimal_scale           ",
		DW_AT_small:                   "DW_AT_small                   ",
		DW_AT_decimal_sign:            "DW_AT_decimal_sign            ",
		DW_AT_digit_count:             "DW_AT_digit_count             ",
		DW_AT_picture_string:          "DW_AT_picture_string          ",
		DW_AT_mutable:                 "DW_AT_mutable                 ",
		DW_AT_threads_scaled:          "DW_AT_threads_scaled          ",
		DW_AT_explicit:                "DW_AT_explicit                ",
		DW_AT_object_pointer:          "DW_AT_object_pointer          ",
		DW_AT_endianity:               "DW_AT_endianity               ",
		DW_AT_elemental:               "DW_AT_elemental               ",
		DW_AT_pure:                    "DW_AT_pure                    ",
		DW_AT_recursive:               "DW_AT_recursive               ",
		DW_AT_signature:               "DW_AT_signature               ",
		DW_AT_main_subprogram:         "DW_AT_main_subprogram         ",
		DW_AT_data_bit_offset:         "DW_AT_data_bit_offset         ",
		DW_AT_const_expr:              "DW_AT_const_expr              ",
		DW_AT_enum_class:              "DW_AT_enum_class              ",
		DW_AT_linkage_name:            "DW_AT_linkage_name            ",
		DW_AT_string_length_bit_size:  "DW_AT_string_length_bit_size  ",
		DW_AT_string_length_byte_size: "DW_AT_string_length_byte_size ",
		DW_AT_rank:                    "DW_AT_rank                    ",
		DW_AT_str_offsets_base:        "DW_AT_str_offsets_base        ",
		DW_AT_addr_base:               "DW_AT_addr_base               ",
		DW_AT_rnglists_base:           "DW_AT_rnglists_base           ",
		DW_AT_dwo_id:                  "DW_AT_dwo_id                  ",
		DW_AT_dwo_name:                "DW_AT_dwo_name                ",
		DW_AT_reference:               "DW_AT_reference               ",
		DW_AT_rvalue_reference:        "DW_AT_rvalue_reference        ",
		DW_AT_macros:                  "DW_AT_macros                  ",
		DW_AT_call_all_calls:          "DW_AT_call_all_calls          ",
		DW_AT_call_all_source_calls:   "DW_AT_call_all_source_calls   ",
		DW_AT_call_all_tail_calls:     "DW_AT_call_all_tail_calls     ",
		DW_AT_call_return_pc:          "DW_AT_call_return_pc          ",
		DW_AT_call_value:              "DW_AT_call_value              ",
		DW_AT_call_origin:             "DW_AT_call_origin             ",
		DW_AT_call_parameter:          "DW_AT_call_parameter          ",
		DW_AT_call_pc:                 "DW_AT_call_pc                 ",
		DW_AT_call_tail_call:          "DW_AT_call_tail_call          ",
		DW_AT_call_target:             "DW_AT_call_target             ",
		DW_AT_call_target_clobbered:   "DW_AT_call_target_clobbered   ",
		DW_AT_call_data_location:      "DW_AT_call_data_location      ",
		DW_AT_call_data_value:         "DW_AT_call_data_value         ",
		DW_AT_noreturn:                "DW_AT_noreturn                ",
		DW_AT_alignment:               "DW_AT_alignment               ",
		DW_AT_export_symbols:          "DW_AT_export_symbols          ",
		DW_AT_deleted:                 "DW_AT_deleted                 ",
		DW_AT_defaulted:               "DW_AT_defaulted               ",
		DW_AT_loclists_base:           "DW_AT_loclists_base           ",
	}
	FormNames = map[int]string{
		DW_FORM_addr: "DW_FORM_addr          ",
		//		DW_FORM_REF:            "DW_FORM_REF           ",
		DW_FORM_block2:         "DW_FORM_block2        ",
		DW_FORM_block4:         "DW_FORM_block4        ",
		DW_FORM_data2:          "DW_FORM_data2         ",
		DW_FORM_data4:          "DW_FORM_data4         ",
		DW_FORM_data8:          "DW_FORM_data8         ",
		DW_FORM_string:         "DW_FORM_string        ",
		DW_FORM_block:          "DW_FORM_block         ",
		DW_FORM_block1:         "DW_FORM_block1        ",
		DW_FORM_data1:          "DW_FORM_data1         ",
		DW_FORM_flag:           "DW_FORM_flag          ",
		DW_FORM_sdata:          "DW_FORM_sdata         ",
		DW_FORM_strp:           "DW_FORM_strp          ",
		DW_FORM_udata:          "DW_FORM_udata         ",
		DW_FORM_ref_addr:       "DW_FORM_ref_addr      ",
		DW_FORM_ref1:           "DW_FORM_ref1          ",
		DW_FORM_ref2:           "DW_FORM_ref2          ",
		DW_FORM_ref4:           "DW_FORM_ref4          ",
		DW_FORM_ref8:           "DW_FORM_ref8          ",
		DW_FORM_ref_udata:      "DW_FORM_ref_udata     ",
		DW_FORM_indirect:       "DW_FORM_indirect      ",
		DW_FORM_sec_offset:     "DW_FORM_sec_offset    ",
		DW_FORM_exprloc:        "DW_FORM_exprloc       ",
		DW_FORM_flag_present:   "DW_FORM_flag_present  ",
		DW_FORM_strx:           "DW_FORM_strx          ",
		DW_FORM_addrx:          "DW_FORM_addrx         ",
		DW_FORM_ref_sup4:       "DW_FORM_ref_sup4      ",
		DW_FORM_strp_sup:       "DW_FORM_strp_sup      ",
		DW_FORM_data16:         "DW_FORM_data16        ",
		DW_FORM_line_strp:      "DW_FORM_line_strp     ",
		DW_FORM_ref_sig8:       "DW_FORM_ref_sig8      ",
		DW_FORM_implicit_const: "DW_FORM_implicit_const",
		DW_FORM_loclistx:       "DW_FORM_loclistx      ",
		DW_FORM_rnglistx:       "DW_FORM_rnglistx      ",
		DW_FORM_ref_sup8:       "DW_FORM_ref_sup8      ",
		DW_FORM_strx1:          "DW_FORM_strx1         ",
		DW_FORM_strx2:          "DW_FORM_strx2         ",
		DW_FORM_strx3:          "DW_FORM_strx3         ",
		DW_FORM_strx4:          "DW_FORM_strx4         ",
		DW_FORM_addrx1:         "DW_FORM_addrx1        ",
		DW_FORM_addrx2:         "DW_FORM_addrx2        ",
		DW_FORM_addrx3:         "DW_FORM_addrx3        ",
		DW_FORM_addrx4:         "DW_FORM_addrx4        ",
		DW_FORM_GNU_addr_index: "DW_FORM_GNU_addr_index",
		DW_FORM_GNU_str_index:  "DW_FORM_GNU_str_index ",
		DW_FORM_GNU_ref_alt:    "DW_FORM_GNU_ref_alt   ",
		DW_FORM_GNU_strp_alt:   "DW_FORM_GNU_strp_alt  ",
	}
	ErrorNames = map[int]string{}
)
